use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();

    let input = read(&mut reader);
    let mut sum: i32 = 0i32;
    for str in input.chars() {
        if "ABC".contains(str) {
            sum += 3;
        } else if "DEF".contains(str) {
            sum += 4;
        } else if "GHI".contains(str) {
            sum += 5;
        } else if "JKL".contains(str) {
            sum += 6;
        } else if "MNO".contains(str) {
            sum += 7;
        } else if "PQRS".contains(str) {
            sum += 8;
        } else if "TUV".contains(str) {
            sum += 9;
        } else if "WXYZ".contains(str) {
            sum += 10;
        }
    }
    println!("{sum}");
}

fn read(reader: &mut BufReader<io::StdinLock>) -> String {
    let mut input = String::new();
    reader.read_line(&mut input).unwrap();
    input
}