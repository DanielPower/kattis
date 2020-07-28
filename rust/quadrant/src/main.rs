use std::io::{self, Read};

fn main() {
    let mut buffer = String::new();
    io::stdin().read_to_string(&mut buffer).unwrap();
    let mut values = buffer.split_whitespace();
    let x = values.next().unwrap().parse::<i16>().unwrap();
    let y = values.next().unwrap().parse::<i16>().unwrap();
    println!("{}", if x > 0 { if y > 0 { 1 } else { 4 } } else { if y > 0 { 2 } else { 3 } });
}
