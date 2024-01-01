use super::checksum;
use super::message::{Byte, Message, MessageType};

/// Serializes all message data into a binary format,
/// added with a checksum.
pub fn serialize(msg: &Message) -> Vec<Byte> {
    let mut binary: Vec<Byte> = Vec::new();

    binary.push(msg.get_head());
    binary.extend(msg.get_payload());

    let checksum: Byte = checksum::generate_checksum(&binary);
    binary.push(checksum);

    binary
}

/// Deserializes a binary format into a message.
/// Some checks are performed to ensure the data is valid.
#[allow(dead_code)]
pub fn deserialize(data: &[Byte]) -> Result<Message, &'static str> {
    // Check if the data is too long.
    if data.len() < 2 {
        return Err("Data is too short");
    }

    // Check if the checksum is correct.
    let body_data: Vec<Byte> = data[..data.len() - 1].to_vec();
    let checksum: Byte = data[data.len() - 1];
    if !checksum::verify_checksum(&body_data, checksum) {
        return Err("Checksum does not match");
    }

    let msg_type = MessageType::from_byte(data[0] >> 6)?;

    let msg_length: Byte = data[0] & 0b0011_1111;
    if msg_length as usize != data.len() {
        return Err("Message length does not match");
    }

    let msg_payload: Vec<Byte> = body_data[1..].to_vec();

    Ok(Message::new(msg_type, msg_length, msg_payload))
}
