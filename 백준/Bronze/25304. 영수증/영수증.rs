use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();

    let sum: i32 = read_single_value(&mut reader);

    let n: i32 = read_single_value(&mut reader);

    let sum2: i32 = (0..n)
        .map(|_| {
            let values = read_multiple_values(&mut reader);
            values[0] * values[1]
        })
        .sum();

    match sum.eq(&sum2) {
        true => println!("Yes"),
        false => println!("No")
    };
}

// 한 줄 입력을 받아 정수로 변환하는 함수
fn read_single_value(reader: &mut BufReader<io::StdinLock>) -> i32 {
    let mut input = String::new();
    reader.read_line(&mut input).expect("Failed to read line");
    input.trim().parse().expect("Failed to parse value")
}

// 한 줄에 있는 여러 값을 읽어서 Vec<i32>로 반환하는 함수
fn read_multiple_values(reader: &mut BufReader<io::StdinLock>) -> Vec<i32> {
    let mut input = String::new();
    reader.read_line(&mut input).expect("Failed to read line");
    input
        .trim()
        .split_whitespace()
        .map(|s| s.parse().expect("Failed to parse number"))
        .collect()
}
