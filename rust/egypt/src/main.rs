use std::io;
use std::io::prelude::BufRead;

fn is_right(a: u32, b: u32, c: u32) {
    if a.pow(2) + b.pow(2) == c.pow(2) {
        println!("right");
    } else {
        println!("wrong");
    }
}

fn main() {
    let stdin = io::stdin();
    for line_result in stdin.lock().lines() {
        let line = line_result.unwrap();
        let mut inputs = line.split_whitespace();
        let a = inputs.next().unwrap().parse::<u32>().unwrap();
        if a == 0 {
            break;
        }
        let b = inputs.next().unwrap().parse::<u32>().unwrap();
        let c = inputs.next().unwrap().parse::<u32>().unwrap();
        if a > b {
            if a > c {
                is_right(b, c, a);
            } else {
                is_right(a, b, c);
            }
        } else if b > c {
            is_right(a, c, b)
        } else {
            is_right(a, b, c);
        }
    }
}
