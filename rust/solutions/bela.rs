use std::io::stdin;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();
    let mut iter = buffer.split_whitespace();
    let num_hands = iter.next().unwrap().parse::<u32>().unwrap();
    let dominant_suite = iter.next().unwrap().parse::<char>().unwrap();
    let mut score = 0;
    for _ in 0..num_hands * 4 {
        buffer = String::new();
        stdin().read_line(&mut buffer).unwrap();
        let mut card = buffer.chars();
        let value = card.next().unwrap();
        let suite = card.next().unwrap();
        score += match value {
            'A' => 11,
            'K' => 4,
            'Q' => 3,
            'J' => {
                if suite == dominant_suite {
                    20
                } else {
                    2
                }
            }
            'T' => 10,
            '9' => {
                if suite == dominant_suite {
                    14
                } else {
                    0
                }
            }
            _ => 0,
        };
    }
    println!("{}", score);
}
