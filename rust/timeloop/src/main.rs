use std::io::{self};

fn main() {
    let mut buffer = String::new();
    io::stdin()
        .read_line(&mut buffer)
        .expect("Failed to read input");
    let count = buffer.trim().parse::<i64>().unwrap();
    for i in 1..count + 1 {
        println!("{} Abracadabra", i);
    }
}
