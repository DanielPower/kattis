use std::io::stdin;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();
    let n = buffer
        .to_string()
        .split_whitespace()
        .next()
        .unwrap()
        .parse::<u32>()
        .unwrap();
    println!("{}", (2 + ((2 as u32).pow(n) - 1)).pow(2));
}
