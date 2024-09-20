use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut chars: Vec<char> = read(&mut reader).chars().collect();
    chars.sort_by(|a, b| b.partial_cmp(a).unwrap());
    let x: String = chars.iter().collect();
    println!("{}", x);
}

fn read(reader: &mut BufReader<io::StdinLock>) -> String {
    let mut input = String::new();
    reader.read_line(&mut input).unwrap();
    input.trim().to_string()
}