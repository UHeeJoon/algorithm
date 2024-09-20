use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    // =====
    let final_string = input.chars()
        .fold(String::new(),
              |mut acc, cur|
                  if acc.len() == 10 {
                      println!("{acc}");
                      cur.to_string()
                  } else {
                      acc.push(cur);
                      acc
                  },
        );
    if !final_string.is_empty() {
        println!("{final_string}");
    }
}
