use std::io::{self, Read};

fn main() -> io::Result<()> {
    let mut buffer = String::new();
    io::stdin().read_to_string(&mut buffer)?;
    let mut initials: Vec<char> = vec![];
    let mut push_next = true;
    for character in buffer.chars() {
        if push_next {
            initials.push(character);
            push_next = false;
        } else if character == '-' {
            push_next = true;
        }
    }
    let initials: String = initials.into_iter().collect();
    println!("{}", initials);
    Ok(())
}
