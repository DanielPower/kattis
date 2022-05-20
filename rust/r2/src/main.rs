use std::io::stdin;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();
    let mut values = buffer.split_whitespace();
    let r1: i16 = values.next().unwrap().parse::<i16>().unwrap();
    let s: i16 = values.next().unwrap().parse::<i16>().unwrap();
    println!("{}", 2 * s - r1);
}
