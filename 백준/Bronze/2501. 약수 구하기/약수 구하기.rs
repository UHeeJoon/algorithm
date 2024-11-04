use std::io::{self, Read};

fn main() {
    let mut reader = io::stdin().lock();
    let mut input: String = String::new();
    reader.read_to_string(&mut input).unwrap();
    let list: Vec<i32> = input.split_whitespace().flat_map(str::parse::<i32>).collect::<Vec<i32>>();
    let mut ans: i32 = 0;
    let mut count: i32 = 0;
    for i in 1..=list[0] {
        if list[0] % i == 0 {
            count += 1;
        }
        if count == list[1] {
            ans = i;
            break;
        }
    }
    if count != list[1] {
        println!("{}", 0);
    } else {
        println!("{}", ans);
    }
}
