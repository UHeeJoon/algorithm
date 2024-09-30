use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    println!("{}", input.trim().chars().collect::<Vec<char>>()[0] as u32 - ('가' as u32) + 1);
}

