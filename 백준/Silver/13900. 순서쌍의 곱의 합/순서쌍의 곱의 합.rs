use std::io::{self, BufRead, BufReader};

fn main() {
    let reader = BufReader::new(io::stdin().lock());
    let mut input = reader.lines();
    let list = input.nth(1).unwrap().unwrap()
        .split_ascii_whitespace()
        .flat_map(str::parse::<i32>)
        .collect::<Vec<i32>>();

    let mut sum: i64 = 0i64;
    let mut multi: i64 = 0i64;
    for window in list.windows(2) {
        sum += window[0] as i64;
        multi += sum * window[1] as i64;
    }
    println!("{}", multi);
}
