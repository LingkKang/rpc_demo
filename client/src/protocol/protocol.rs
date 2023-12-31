use super::checksum;

/// Maximum size of a protocol message.
/// It is counted in bytes.
const MAX_PROTOCOL_SIZE: usize = 64;

/// Maximum size of payload
/// comes from `MAX_PROTOCOL_SIZE - 2`.
const MAX_PAYLOAD_SIZE: usize = MAX_PROTOCOL_SIZE - 2;

/// The type (action) of a message which is the first two bits of a message.
///
/// | Bits | Type     |
/// | :--: | :------- |
/// | `00` | Error    |
/// | `01` | Request  |
/// | `10` | Response |
///
/// Note that `11` is reserved.
pub enum MessageType {
    /// Message type for error,
    /// maps to `0b00` in the first two bits.
    ERROR,

    /// Message type for request,
    /// maps to `0b01` in the first two bits.
    REQUEST,

    /// Message type for response,
    /// maps to `0b10` in the first two bits.
    RESPONSE,
}

impl MessageType {
    /// Convert a `MessageType` to a bit (`u8`) as defined in the protocol.
    fn to_bits(&self) -> u8 {
        match self {
            MessageType::ERROR => 0b00,
            MessageType::REQUEST => 0b01,
            MessageType::RESPONSE => 0b10,
        }
    }
}

/// The protocol message.
/// 
/// A protocol message is defined as follows:
/// 
/// | component     | size (bits)  | description |
/// | :------------ | :----------: | :---------- |
/// | `msg_type`    | 2            | The type (action) of the message |
/// | `msg_length`  | 6            | The length of the entire message |
/// | `msg_payload` | 0 - 62 bytes | The payload of the message       |
/// | checksum      | 8            | The checksum of the message      |
/// 
/// Note that:
/// 1. the `msg_type` and `msg_length` form the first byte of the message,
/// which is the first two bits of the byte for `msg_type` and the following six bits for `msg_length`.
/// 2. the value of `msg_length` is the length of the entire message in bytes,
/// so the whole protocol message is capable of carrying `2^6 - 2 = 62` bytes of payload.
/// 3. the checksum is not included in the `Message` struct,
/// it is calculated and appended when serializing the message.
pub struct Message {
    /// The type (action) of the message,
    /// which will be encoded in the first two bits of the message.
    msg_type: MessageType,

    /// The length of the entire message,
    /// which will be encoded in the following 6 bits of the message.
    msg_length: u8,

    /// The payload of the message,
    /// should always be a multiple of 8 bits (1 byte).
    msg_payload: Vec<u8>,
}
