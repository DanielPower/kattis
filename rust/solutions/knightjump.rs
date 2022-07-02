use std::collections::VecDeque;
use std::io::{self, Read};
use std::process;

const KNIGHT_MOVES: [(i64, i64); 8] = [
    (-2, -1),
    (-2, 1),
    (2, -1),
    (2, 1),
    (-1, -2),
    (-1, 2),
    (1, -2),
    (1, 2),
];

fn to_1d(x: i64, y: i64, s: i64) -> i64 {
    y * s + x
}

fn to_2d(i: i64, s: i64) -> (i64, i64) {
    (i % s, i / s)
}

fn is_valid(nodes: &Vec<bool>, x: i64, y: i64, i: i64, s: i64) -> bool {
    x >= 0 && y >= 0 && x < s && y < s && nodes[i as usize]
}

fn get_next_nodes(nodes: &mut Vec<bool>, i: i64, s: i64) -> Vec<i64> {
    let (x, y) = to_2d(i, s);
    let mut next_nodes: Vec<i64> = vec![];
    for (dx, dy) in KNIGHT_MOVES.iter() {
        let candidate_node = to_1d(x + dx, y + dy, s);
        if is_valid(&nodes, x + dx, y + dy, candidate_node, s) {
            nodes[candidate_node as usize] = false;
            next_nodes.push(candidate_node)
        }
    }
    next_nodes
}

fn main() {
    let mut buffer = String::new();

    io::stdin()
        .read_line(&mut buffer)
        .expect("Unable to read grid size input");
    let size = buffer
        .to_string()
        .trim()
        .parse::<i64>()
        .expect("Grid size input could not be parsed as integer");

    io::stdin()
        .read_to_string(&mut buffer)
        .expect("Unable to read input");

    let mut nodes: Vec<bool> = vec![];
    let mut read_index: i64 = 0;
    let mut start_index: i64 = 0;

    for character in buffer.chars() {
        match character {
            '.' => {
                nodes.push(true);
                read_index += 1;
            }
            'K' => {
                if read_index == 0 {
                    println!("0");
                    process::exit(0x0100);
                }
                nodes.push(false);
                start_index = read_index;
                read_index += 1;
            }
            '#' => {
                nodes.push(false);
                read_index += 1;
            }
            _ => (),
        }
    }

    let mut open: VecDeque<(i64, i64)> = vec![].into_iter().collect();
    let next_nodes = get_next_nodes(&mut nodes, start_index, size); 

    for &index in next_nodes.iter() {
        open.push_back((index, 1));
    }

    while !open.is_empty() {
        let (index, depth) = open.pop_front().expect("There should be a node here");

        if index == 0 {
            println!("{}", depth);
            process::exit(0x0100);
        }

        let next_nodes = get_next_nodes(&mut nodes, index, size);
        for next_index in next_nodes.iter() {
            open.push_back((*next_index, depth + 1));
        }
    }
    println!("-1");
}
