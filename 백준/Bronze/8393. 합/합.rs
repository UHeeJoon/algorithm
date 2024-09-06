use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();

    let mut input = String::new();
    reader.read_line(&mut input).expect("");

    let num: i32 = input.trim().parse().expect("");

    (1..=num).reduce(|a,b|  a + b).and_then(|x| Some(println!("{x}")));
}
