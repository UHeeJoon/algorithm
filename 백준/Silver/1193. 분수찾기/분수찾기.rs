use std::io::{self, Read};

fn main() {
    let mut reader = io::stdin().lock();
    let mut input: String = String::new();
    reader.read_to_string(&mut input).unwrap();
    let mut num = input.trim().parse::<i32>().unwrap();
    let mut i: i32 = 1i32;
    while num > 0 {
        num -= i;
        i += 1;
    }
    i -= 1;
    num += i;
    let (numerator, denominator) = if i % 2 == 0 {
        (num, i + 1 - num)
    } else {
        (i + 1 - num, num)
    };

    println!("{}/{}", numerator, denominator);
}
