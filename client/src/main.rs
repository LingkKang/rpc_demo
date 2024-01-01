mod protocol;

use std::{
    io::{Read, Write},
    net::{Shutdown, TcpStream},
};

use rand::Rng;

use crate::protocol::{
    message::{Byte, Message},
    protocol::serialize,
};

const URL: &str = "test.lingkang.dev:8333";
const MAX_TASKS: usize = 4;

fn main() {
    let mut stream = match TcpStream::connect(URL) {
        Ok(stream) => stream,
        Err(e) => {
            println!("{}", e);
            return;
        }
    };
    receive_message(&mut stream);

    for _ in 0..MAX_TASKS {
        let msg = Message::new_request(collect_a_task());
        let msg_str: Vec<Byte> = serialize(&msg);
        print!("Sending (in decimal): {:?}\n", msg_str);
        stream.write_all(&msg_str).unwrap();
        receive_message(&mut stream);
    }

    stream.shutdown(Shutdown::Both).unwrap();

    println!("Exiting...");
}

/// Receives a message from the server and prints it.
fn receive_message(stream: &mut TcpStream) {
    let mut buffer: [Byte; 64] = [0; 64];
    let bytes_read = stream.read(&mut buffer).unwrap();
    println!(
        "Received: {}",
        String::from_utf8_lossy(&buffer[..bytes_read])
    );
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
    println!("a = {}, b = {}", a, b);
    let mut sides: Vec<Byte> = Vec::new();
    sides.extend(a.to_be_bytes().to_vec());
    sides.extend(b.to_be_bytes().to_vec());
    sides
}
