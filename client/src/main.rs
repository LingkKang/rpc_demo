mod logger;
mod protocol;

use std::{
    io::{Read, Write},
    net::{Shutdown, TcpStream},
};

use rand::Rng;

use crate::protocol::{
    message::{Byte, Message, MessageType, MAX_PROTOCOL_SIZE},
    protocol::{deserialize, serialize},
};

use crate::logger::logger::Logger;

const URL: &str = "test.lingkang.dev:8333";
const MAX_TASKS: usize = 4;

#[allow(unreachable_code)]
fn main() {
    Logger::init(Some(log::LevelFilter::Debug));

    let mut stream = match TcpStream::connect(URL) {
        Ok(stream) => stream,
        Err(e) => {
            log::error!("{}", e);
            return;
        }
    };

    for _ in 0..MAX_TASKS {
        let msg = Message::new_request(collect_a_task());
        let msg_str: Vec<Byte> = serialize(&msg);
        log::debug!("Sending (in hex): {}", bytes_to_hex_str(&msg_str));
        stream.write_all(&msg_str).unwrap();
        process_message(receive_message(&mut stream));
    }

    stream.shutdown(Shutdown::Both).unwrap();

    log::info!("Exiting...");
}

/// Receives a message from the server and parse it.
fn receive_message(stream: &mut TcpStream) -> Message {
    let mut buffer: [Byte; MAX_PROTOCOL_SIZE] = [0; MAX_PROTOCOL_SIZE];
    let bytes_read = stream.read(&mut buffer).unwrap();

    deserialize(&buffer[..bytes_read]).unwrap()
}

/// Generates a random [`f64`] number.
fn get_random_f64() -> f64 {
    let mut rng = rand::thread_rng();
    rng.gen::<f64>() * 100.0
}

/// Generates a pair of random [`f64`] numbers,
/// which will be used as the sides of a right triangle.
fn get_sides() -> (f64, f64) {
    (get_random_f64(), get_random_f64())
}

/// Basically generates a task of calculating the
/// hypotenuse of a right triangle,
/// and returns the sides of the triangle as a [`Vec`] of [`Byte`]s.
fn collect_a_task() -> Vec<Byte> {
    let (a, b) = get_sides();
    log::debug!("a = {}, b = {}", a, b);
    let mut sides: Vec<Byte> = Vec::new();
    sides.extend(a.to_be_bytes().to_vec());
    sides.extend(b.to_be_bytes().to_vec());
    sides
}

/// Process a message, basically check the type of the message
/// and call the corresponding function.
fn process_message(msg: Message) {
    match msg.get_type() {
        MessageType::ERROR => todo!(),
        MessageType::REQUEST => todo!(),
        MessageType::RESPONSE => process_response(msg),
    }
}

/// Process a message of type [`MessageType::RESPONSE`].
fn process_response(msg: Message) {
    let payload: &Vec<Byte> = msg.get_payload();
    // Convert the vector to an array.
    let mut arr: [Byte; 8] = [0; 8];
    arr.copy_from_slice(payload);
    let hypotenuse: f64 = f64::from_be_bytes(arr);
    log::info!("Received hypotenuse: {hypotenuse}");
}

/// Converts a [`Vec`] of [`Byte`]s to a [`String`] of hexadecimal numbers.
fn bytes_to_hex_str(bytes: &[Byte]) -> String {
    bytes.iter().map(|byte| format!("{byte:X}")).collect()
}
