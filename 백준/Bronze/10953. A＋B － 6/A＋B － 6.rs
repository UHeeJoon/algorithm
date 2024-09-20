use std::io::{self, BufRead, BufReader, Write};
fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    // =====

    let num: i32 = input.trim().parse().unwrap();
    (0..num).for_each(|_| {
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();
        
        let split: Vec<&str> = input.split(",").collect();
        println!("{}",
             split[0].trim().parse::<i32>().unwrap() + split[1].trim().parse::<i32>().unwrap()
        );
    })
}
