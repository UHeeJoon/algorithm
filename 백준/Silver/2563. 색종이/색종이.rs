use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let mut list: [[bool; 101]; 101] = [[false; 101]; 101];

    for _ in 0..input.trim().parse().unwrap() {
        input.clear();
        reader.read_line(&mut input).unwrap();
        let input_list: Vec<usize> = input.split_whitespace().flat_map(str::parse::<usize>).collect::<Vec<usize>>();


        for i in input_list[0]..(input_list[0] + 10) {
            for j in input_list[1]..(input_list[1] + 10) {
                list[i][j] = true;
            }
        }
    }

    let count: usize = list.iter()
            .flat_map(|&row| row)
            .filter(|&cell| cell)
            .count();
        
    println!("{}", count);
}
