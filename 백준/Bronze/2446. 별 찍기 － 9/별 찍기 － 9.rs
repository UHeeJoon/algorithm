use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let input = input.trim().parse().unwrap();
    for i in (1..=input).rev().chain(2..=input) {
        println!("{}{}", " ".repeat(input - i), "*".repeat(2 * i - 1));
    }
}
