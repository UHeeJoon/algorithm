use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();

    let mut input = String::new();
    reader.read_line(&mut input).expect("");

    let mut numbers = input.trim().split_whitespace();
    let num1: i32 = numbers.next().expect("").parse().expect("");
    let num2: i32 = numbers.next().expect("").parse().expect("");

    let mut answer = "==";

    if num1 < num2 {
        answer = "<";
    } else if num1 > num2 {
        answer = ">";
    }

    println!("{answer}");
}
