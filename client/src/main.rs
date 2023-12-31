use std::{
    io::{Read, Write},
    net::TcpStream,
};

fn main() {
    let url = "test.lingkang.dev:8333";
    let mut stream = TcpStream::connect(url).unwrap();
    receive_message(&mut stream);

    let greeting = "Hello from Rust!\n";
    print!("Sending: {}", greeting);
    let _ = stream.write_all(greeting.as_bytes());
    receive_message(&mut stream);

    let exit = "exit\n";
    print!("Sending: {}", exit);
    let _ = stream.write_all(exit.as_bytes());
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
