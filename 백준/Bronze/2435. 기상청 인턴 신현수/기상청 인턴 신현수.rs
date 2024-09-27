use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut n_k: String = String::new();
    reader.read_line(&mut n_k).unwrap();
    let k: usize = n_k.split_whitespace().flat_map(str::parse::<usize>).collect::<Vec<_>>()[1];
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let list: Vec<i32> = input.split_whitespace().flat_map(str::parse::<i32>).collect();
    let mut max: i32 = i32::MIN;
    (0..=list.len() - k).for_each(|i| {
        max = max.max(list[i..i + k].iter().sum());
    });
    println!("{max}");
}

