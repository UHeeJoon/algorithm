use std::cmp::min;
use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    // =====

    for i in (0..input.len()).step_by(10)  {
        println!("{}", &input[i..min(i + 10, input.len())]);
    }
}
