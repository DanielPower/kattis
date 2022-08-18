use std::io;

fn main() {
    let mut buffer = String::new();
    io::stdin()
        .read_line(&mut buffer)
        .expect("Error reading input");
    let mut carrots_iter = buffer.split_whitespace();
    carrots_iter.next();
    let carrots = carrots_iter
        .next()
        .expect("Failed to get carrot number")
        .trim()
        .parse::<i64>()
        .unwrap();
    println!("{}", carrots);
}
