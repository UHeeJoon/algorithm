use std::cmp::max;
use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut t: String = String::new();
    reader.read_line(&mut t).unwrap();

    for _ in 0..t.trim().parse::<i32>().unwrap() {
        reader.read_line(&mut String::new()).unwrap();
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();

        let mut list: Vec<i32> = input.split_whitespace().flat_map(str::parse).collect();
        list.insert(0, 0);

        let size: usize = list.iter().len();
        (1..size).for_each(|i| list[i] += list[i - 1]);

        let mut result: i32 = -2_000_000_000i32;
        (1..size).for_each(|i|
            (0..i).for_each(|j| result = max(list[i] - list[j], result))
        );
        println!("{result}");
    }
}