use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut n: String = String::new();
    reader.read_line(&mut n).unwrap();

    let n = n.trim().parse().unwrap();
    let mut count: i32 = 0i32;

    for _ in 0..n {
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();

        let mut chs: [bool; 26] = [false; 26];
        let mut prev: Option<char> = None;
        let mut flag: bool = true;
        for c in input.trim().chars() {
            if Some(c) != prev {
                let idx: usize = c as usize - 97;
                if chs[idx] {
                    flag = false;
                    break;
                }
                chs[idx] = true;
            }
            prev = Some(c);
        }
        if flag {
            count += 1;
        }
    }
    println!("{}", count);
}
