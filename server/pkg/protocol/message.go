package protocol

import (
	"encoding/binary"
	"fmt"
	"math"
)

type MessageCode byte

const (
	ERROR    MessageCode = 0b00
	REQUEST  MessageCode = 0b01
	RESPONSE MessageCode = 0b10
)

type Message struct {
	msg_code     MessageCode
	msg_len      byte
	msg_payload  []byte
	msg_checksum byte
}

const MSG_HEADER_SIZE = 1
const MSG_CHECKSUM_SIZE = 1

func GetPayloadLength(head byte) byte {
	return ((head << 2) >> 2) - MSG_HEADER_SIZE - MSG_CHECKSUM_SIZE
}

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

func ParsePayloadToFloat(payload []byte) ([]float64, error) {
	if len(payload)%8 != 0 {
		return nil, fmt.Errorf(
			"payload length is not a multiple of 8 when handling a REQUEST")
	}
	var floats []float64
	for i := 0; i < len(payload); i += 8 {
		eight_bits := binary.BigEndian.Uint64(payload[i : i+8])
		floats = append(floats, math.Float64frombits(eight_bits))
	}
	return floats, nil
}
