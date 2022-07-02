use std::io::stdin;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();
    let mut iter = buffer.split_whitespace();
    let x: u8 = iter.next().unwrap().parse::<u8>().unwrap();
    let y: u8 = iter.next().unwrap().parse::<u8>().unwrap();
    let n: u8 = iter.next().unwrap().parse::<u8>().unwrap();

    for i in 1..n + 1 {
        let mut f: bool = false;
        if i % x == 0 {
            print!("Fizz");
            f = true;
        }
        if i % y == 0 {
            print!("Buzz");
            f = true;
        }
        if !f {
            print!("{}", i);
        }
        println!();
    }
}
