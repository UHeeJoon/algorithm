use std::io::{self, BufRead, BufReader, Write};
fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    // =====

    let mut list: [usize; 26] = [0; 26];
    input.chars().for_each(|c| {
        if c.is_ascii_lowercase() {
            list[c as usize - 97] += 1;
        }
    });
    let x: String = list.iter().map(|i| i.to_string() + " ").collect();
    println!("{}", x);
}
