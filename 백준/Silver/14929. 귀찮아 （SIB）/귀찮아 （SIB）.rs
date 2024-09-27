use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut n_k: String = String::new();
    reader.read_line(&mut n_k).unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let list: Vec<i64> = input.split_whitespace().flat_map(str::parse::<i64>).collect();
    let mut sum: i64 = list.iter().sum();
    let result: i64 = list.iter().fold(0, |acc, v| {sum -= v; acc + v * sum});
    println!("{result}");
}

