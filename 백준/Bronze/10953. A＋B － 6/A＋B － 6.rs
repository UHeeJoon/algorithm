use std::io::{self, BufRead, BufReader, Write};
fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    // =====

    let num: i32 = input.trim().parse().unwrap();
    (0..num).for_each(|_| {
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();

        let sum: i32 = input.split(",").map(str::trim).flat_map(str::parse::<i32>).sum::<i32>();
        println!("{}", sum);
    })
}
