use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let mut change: i32 = 1000 - input.trim().parse::<i32>().unwrap();
    let coins = [500, 100, 50, 10, 5, 1];
    let mut answer = 0;

    for &coin in &coins {
        answer += change / coin;
        change %= coin;
    }

    println!("{}", answer);
}