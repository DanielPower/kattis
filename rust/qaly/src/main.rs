use std::io::{self, Read};

fn main() {
    let mut buffer = String::new();
    io::stdin()
        .read_to_string(&mut buffer)
        .expect("Couldn't read input");
    let mut total_quality: f64 = 0.0;
    let mut lines = buffer.lines();
    lines.next();
    for line in lines {
        let mut iter = line.split_whitespace();
        let quality = iter
            .next()
            .expect("Couldn't read input")
            .parse::<f64>()
            .expect("Expected input to be floating point number");
        let quantity = iter
            .next()
            .expect("Couldn't read input")
            .parse::<f64>()
            .expect("Expected input to be floating point number");
        total_quality += quality * quantity;
    }
    println!("{}", total_quality);
}
