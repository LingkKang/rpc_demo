mod protocol;

use std::{
    io::{Read, Write},
    net::TcpStream,
    vec,
};

use crate::protocol::{message::Message, protocol::serialize};

fn main() {
    let url = "test.lingkang.dev:8333";
    let mut stream = TcpStream::connect(url).unwrap();
    receive_message(&mut stream);

    let msg = Message::new_request(vec![0x00, 0x02]);
    let msg_str = serialize(&msg);
    print!("Sending (in decimal): {:?}\n", msg_str);
    let _ = stream.write_all(&msg_str);
    receive_message(&mut stream);

    println!("Exiting...");
}

fn receive_message(stream: &mut TcpStream) {
    let mut buffer = [0; 1024];
    let bytes_read = stream.read(&mut buffer).unwrap();
    println!(
        "Received: {}",
        String::from_utf8_lossy(&buffer[..bytes_read])
    );
}
