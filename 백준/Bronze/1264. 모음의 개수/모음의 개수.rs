use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    loop  {
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();
        if input.trim().eq("#") {
            break;
        }
        let count: i32 = input.trim().to_uppercase().split("").filter(|&c| c.eq("A") || c.eq("E") || c.eq("I") || c.eq("O") || c.eq("U")).count() as i32;
        println!("{}", count);
    }
}

