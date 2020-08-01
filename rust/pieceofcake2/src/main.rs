use std::cmp::max;
use std::io::stdin;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();
    let mut inputs = buffer.split_whitespace();
    let n = inputs.next().unwrap().parse::<u32>().unwrap();
    let x1 = inputs.next().unwrap().parse::<u32>().unwrap();
    let x2 = n - x1;
    let y1 = inputs.next().unwrap().parse::<u32>().unwrap();
    let y2 = n - y1;

    println!("{}", max(max(max(x1 * y1, x1 * y2), x2 * y1), x2 * y2) * 4);
}
