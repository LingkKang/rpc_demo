package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// Enumeration of MessageCode,
// which will indicate the action/information type of the message.
type MessageCode byte

const (
	ERROR    MessageCode = 0b00
	REQUEST  MessageCode = 0b01
	RESPONSE MessageCode = 0b10
)

// A self-defined protocol message.
type Message struct {
	msg_code     MessageCode
	msg_len      byte
	msg_payload  []byte
	msg_checksum byte
}

const MSG_HEADER_SIZE = 1
const MSG_CHECKSUM_SIZE = 1

// Given the header byte,
// return the payload length,
// which is the last 6 bits of the header.
func GetPayloadLength(head byte) byte {
	return ((head << 2) >> 2) - MSG_HEADER_SIZE - MSG_CHECKSUM_SIZE
}

// Given the processed byte of `MessageCode`,
// return the corresponding enum.
func convertMessageCodeFromByte(code byte) MessageCode {
	switch code {
	case 0b00:
		return ERROR
	case 0b01:
		return REQUEST
	case 0b10:
		return RESPONSE
	default:
		panic("Unkown message code")
	}
}

// Deserialize bytes of message to an acutal `Message`.
func NewMessageFromBytes(
	head byte,
	payload []byte,
	checksum byte) (Message, error) {
	msg_code := convertMessageCodeFromByte(head >> 6)
	msg_len := GetPayloadLength(head)

	// Validate checksum.
	data := append([]byte{head}, payload...)
	if !validateChecksum(data, checksum) {
		return Message{}, fmt.Errorf("checksum validation failed")
	}

	return Message{msg_code, msg_len, payload, checksum}, nil
}

func GetMessageCode(msg Message) MessageCode {
	return msg.msg_code
}

func GetMessagePayload(msg Message) []byte {
	return msg.msg_payload
}

// Parse the big-endian bytes of payload to actual `float64`s.
func ParsePayloadToFloat64s(payload []byte) ([]float64, error) {
	if len(payload)%8 != 0 {
		return nil, fmt.Errorf(
			"payload length is not a multiple of 8 when handling a REQUEST")
	}
	var floats []float64
	for i := 0; i < len(payload); i += 8 {
		eight_bits := binary.BigEndian.Uint64(payload[i : i+8])
		f := math.Float64frombits(eight_bits)
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return nil, fmt.Errorf("encountered Nan or Inf in payload")
		}
		floats = append(floats, f)
	}
	return floats, nil
}

// Generate the first byte (head) of the message.
func GetMessageHead(msg Message) byte {
	return (byte(msg.msg_code) << 6) | msg.msg_len
}

// Generate a new message in the type of `RESPONSE` as server feedback.
func NewResponseMessage(result float64) Message {
	// Convert the passed in reslut float to bytes.
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, result)

	// Construct basic `Message` structure.
	msg := Message{
		msg_code:    RESPONSE,
		msg_len:     8 + MSG_HEADER_SIZE + MSG_CHECKSUM_SIZE,
		msg_payload: buf.Bytes(),
	}

	// Calculate and fill in the checksum.
	data := append([]byte{GetMessageHead(msg)}, msg.msg_payload...)
	msg.msg_checksum = generateChecksum(data)

	return msg
}

// Convert a `Message` to a slice of bytes.
func SerializeMessage(msg Message) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, GetMessageHead(msg))
	binary.Write(buf, binary.BigEndian, msg.msg_payload)
	binary.Write(buf, binary.BigEndian, msg.msg_checksum)
	return buf.Bytes()
}
