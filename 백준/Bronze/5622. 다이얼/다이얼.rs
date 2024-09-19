use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();

    let input = read(&mut reader);
    let sum: i32 = input.chars()
        .map(|c| match c {
            'A' | 'B' | 'C' => 3,
            'D' | 'E' | 'F' => 4,
            'G' | 'H' | 'I' => 5,
            'J' | 'K' | 'L' => 6,
            'M' | 'N' | 'O' => 7,
            'P' | 'Q' | 'R' | 'S' => 8,
            'T' | 'U' | 'V' => 9,
            'W' | 'X' | 'Y' | 'Z' => 10,
            _ => 0, // 다른 문자는 0점 처리
        })
        .sum();
    println!("{sum}");
}

fn read(reader: &mut BufReader<io::StdinLock>) -> String {
    let mut input = String::new();
    reader.read_line(&mut input).unwrap();
    input
}