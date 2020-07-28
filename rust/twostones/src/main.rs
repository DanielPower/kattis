use std::io::{self};

fn main() {
    let mut buffer = String::new();
    io::stdin().read_line(&mut buffer).expect("Failed");
    let stone_count = buffer.trim().parse::<i64>().unwrap();
    println!("{}", if stone_count % 2 == 0 {"Bob"} else {"Alice"});
}
