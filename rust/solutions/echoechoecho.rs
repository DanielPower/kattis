use std::io;

fn main() {
    let mut buffer = String::new();
    io::stdin()
        .read_line(&mut buffer)
        .expect("Error reading input");
    buffer = buffer.trim().to_string();
    println!("{0} {0} {0}", buffer);
}
