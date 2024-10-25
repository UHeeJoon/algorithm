use std::io::{self, BufRead, BufReader};

fn main() {
    let reader = BufReader::new(io::stdin().lock());
    let mut lines = reader.lines();
    let (r1, s) = {
        let nums = lines.next().unwrap().unwrap();
        let mut iter = nums.split_ascii_whitespace().map(str::parse::<i32>);
        (iter.next().unwrap().unwrap(), iter.next().unwrap().unwrap())
    };
    println!("{}", (s * 2) - r1);
}
