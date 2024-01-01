//! A protocol message is defined as follows:
//!
//! | component     | size (bits)  | description                      |
//! | :------------ | :----------: | :------------------------------- |
//! | `msg_type`    | 2            | The type (action) of the message |
//! | `msg_length`  | 6            | The length of the entire message |
//! | `msg_payload` | 0 - 62 bytes | The payload of the message       |
//! | checksum      | 8            | The checksum of the message      |
//!
//! Note that:
//! 1. the `msg_type` and `msg_length` form the first byte of the message,
//! which is the first two bits of the byte for `msg_type` and the following six bits for `msg_length`.
//! 2. the value of `msg_length` is the length of the entire message in bytes,
//! so the whole protocol message is capable of carrying `2^6 - 2 = 62` bytes of payload.
//! 3. the checksum is not included in the [`Message`] struct,
//! it is calculated and appended when serializing the message.

/// A byte is logically equivalent to an 8-bit unsigned integer.
pub type Byte = u8;

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
///
// #[derive(Debug, PartialEq)]
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
    /// Convert a `MessageType` to a [`Byte`] as defined in the protocol.
    fn to_byte(&self) -> Byte {
        match self {
            MessageType::ERROR => 0b00,
            MessageType::REQUEST => 0b01,
            MessageType::RESPONSE => 0b10,
        }
    }

    pub fn from_byte(byte: Byte) -> Result<MessageType, &'static str> {
        match byte {
            0b00 => Ok(MessageType::ERROR),
            0b01 => Ok(MessageType::REQUEST),
            0b10 => Ok(MessageType::RESPONSE),
            _ => Err("Invalid message type"),
        }
    }
}

/// A protocol message.
/// See the module-level documentation for more details.
pub struct Message {
    /// The type (action) of the message,
    /// which will be encoded in the first two bits of the message.
    msg_type: MessageType,

    /// The length of the entire message,
    /// which will be encoded in the following 6 bits of the message.
    msg_length: Byte,

    /// The payload of the message,
    /// should always be a multiple of 8 bits (1 byte).
    msg_payload: Vec<Byte>,
}

impl Message {
    /// Create a new message.
    pub fn new(msg_type: MessageType, msg_length: Byte, msg_payload: Vec<Byte>) -> Message {
        assert!(msg_length as usize <= MAX_PAYLOAD_SIZE);
        Message {
            msg_type,
            msg_length,
            msg_payload,
        }
    }

    pub fn new_request(msg_payload: Vec<Byte>) -> Message {
        Message::new(
            MessageType::REQUEST,
            msg_payload.len() as Byte + 2,
            msg_payload,
        )
    }

    /// Get the first byte of the message,
    /// which is the first two bits for `msg_type` and the following six bits for `msg_length`.
    pub fn get_head(&self) -> Byte {
        (self.msg_type.to_byte() << 6) | self.msg_length
    }

    /// Get the payload of the message.
    pub fn get_payload(&self) -> &Vec<Byte> {
        &self.msg_payload
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_message_type_to_byte() {
        assert_eq!(MessageType::ERROR.to_byte(), 0b00);
        assert_eq!(MessageType::REQUEST.to_byte(), 0b01);
        assert_eq!(MessageType::RESPONSE.to_byte(), 0b10);
    }

    #[test]
    fn test_message_type_from_byte() {
        assert!(matches!(
            MessageType::from_byte(0b00).unwrap(),
            MessageType::ERROR
        ));

        assert!(matches!(
            MessageType::from_byte(0b01).unwrap(),
            MessageType::REQUEST
        ));

        assert!(matches!(
            MessageType::from_byte(0b10).unwrap(),
            MessageType::RESPONSE
        ));

        assert!(MessageType::from_byte(0b11).is_err());
    }

    #[test]
    fn test_message_new() {
        let msg_type = MessageType::REQUEST;
        let msg_length = 0b0000_0011;
        let msg_payload = vec![0x01];
        let payload_cpy = msg_payload.clone();

        let msg = Message::new(msg_type, msg_length, msg_payload);

        assert!(matches!(msg.msg_type, MessageType::REQUEST));
        assert_eq!(msg.msg_length, msg_length);
        assert_eq!(msg.msg_payload, payload_cpy);
    }

    #[test]
    fn test_message_get_head() {
        let msg_type = MessageType::REQUEST;
        let msg_length: Byte = 0b0000_0100;
        let msg_payload: Vec<Byte> = vec![0x01, 0x02];
        let payload_cpy: Vec<Byte> = msg_payload.clone();

        let msg = Message::new(msg_type, msg_length, msg_payload);
        assert_eq!(msg.get_head(), 0b0100_0100);

        let msg_type = MessageType::RESPONSE;
        let msg_length: Byte = 0b0000_0011;

        let msg = Message::new(msg_type, msg_length, payload_cpy);
        assert_eq!(msg.get_head(), 0b1000_0011);
    }

    #[test]
    fn test_message_get_payload() {
        let msg_type = MessageType::REQUEST;
        let msg_length: Byte = 0b0000_0100;
        let msg_payload: Vec<Byte> = vec![0x01, 0x02];
        let msg = Message::new(msg_type, msg_length, msg_payload);
        assert_eq!(msg.get_payload(), &vec![0x01, 0x02]);

        let msg_type = MessageType::RESPONSE;
        let msg_length: Byte = 0b0000_0011;
        let msg_payload: Vec<Byte> = vec![0x01];
        let msg = Message::new(msg_type, msg_length, msg_payload);
        assert_eq!(msg.get_payload(), &vec![0x01]);
    }
}
