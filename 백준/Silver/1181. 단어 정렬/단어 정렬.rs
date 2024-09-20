use std::io::{self, BufRead, BufReader, Write};
fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    let mut num: String = String::new();
    reader.read_line(&mut num).unwrap();
    // =====

    let num = num.trim().parse().unwrap();
    let mut list: Vec<String> = (0..num).map(|_| {
        input.clear();
        reader.read_line(&mut input).unwrap();
        input.clone().trim().to_string()
    }).collect();

    let cmp = |a: &String, b: &String| {
        if a.len() != b.len() {
            a.len().cmp(&b.len())
        } else {
            a.chars().cmp(b.chars())
        }
    };
    list.sort_by(cmp);
    list.dedup();
    list.iter().for_each(|l| println!("{l}"));
}