use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let input: Vec<String>  = (0..5).map(|_| read(&mut reader)).collect();
    let map: Vec<_> = (0..15).flat_map(|i| input.iter().filter_map(move |str| str.get(i..=i))).collect();
    println!("{}", map.join(""));
}

fn read(reader: &mut BufReader<io::StdinLock>) -> String {
    let mut input = String::new();
    reader.read_line(&mut input).unwrap();
    input.trim().to_string()
}