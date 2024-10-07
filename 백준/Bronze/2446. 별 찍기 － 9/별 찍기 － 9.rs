use std::fmt::Write;
use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut output: String = String::new();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let input = input.trim().parse().unwrap();
    for i in (1..=input).rev().chain(2..=input) {
        writeln!(output, "{}{}", " ".repeat(input - i), "*".repeat(2 * i - 1)).ok();
    }
    println!("{output}");
}
