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
        input.to_owned().trim().to_string()
    }).collect();

    list.sort_unstable_by_key(|key| (key.len(), key.clone()));
    list.dedup();
    println!("{}", list.join("\n"));
}
