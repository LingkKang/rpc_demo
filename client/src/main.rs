mod protocol;

use std::{
    io::{Read, Write},
    net::TcpStream,
    vec,
};

use crate::protocol::{
    message::{Byte, Message},
    protocol::serialize,
};

const URL: &str = "test.lingkang.dev:8333";

fn main() {
    let mut stream = match TcpStream::connect(URL) {
        Ok(stream) => stream,
        Err(e) => {
            println!("{}", e);
            return;
        }
    };
    receive_message(&mut stream);

    let msg = Message::new_request(vec![0x00, 0x02]);
    let msg_str: Vec<Byte> = serialize(&msg);
    print!("Sending (in decimal): {:?}\n", msg_str);
    let _ = stream.write_all(&msg_str);
    receive_message(&mut stream);

    println!("Exiting...");
}

fn receive_message(stream: &mut TcpStream) {
    let mut buffer: [Byte; 64] = [0; 64];
    let bytes_read = stream.read(&mut buffer).unwrap();
    println!(
        "Received: {}",
        String::from_utf8_lossy(&buffer[..bytes_read])
    );
}
