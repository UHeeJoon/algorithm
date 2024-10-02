use std::io::{self, BufRead, BufReader};

fn main() {
    let reader = BufReader::new(io::stdin().lock());
    for line in reader.lines() {
        let nums: Vec<i32> = line.unwrap()
            .split_whitespace()
            .filter_map(|x| x.parse::<i32>().ok())
            .collect();

        let (a, b): (i32, i32) = (nums[0], nums[1]);

        if a == 0 && b == 0 {
            break;
        }

        let result = match (a % b == 0, b % a == 0) {
            (true, false) => "multiple",
            (false, true) => "factor",
            _ => "neither",
        };
        println!("{}", result);
    }
}

