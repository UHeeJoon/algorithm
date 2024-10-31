use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let x_n: Vec<usize> = input.split_whitespace().flat_map(str::parse::<usize>).collect();
    input.clear();
    reader.read_line(&mut input).unwrap();
    let list: Vec<i32> = input.split_whitespace().flat_map(str::parse::<i32>).collect();

    let mut sum: i32 = list[0..x_n[1]].iter().sum();
    let mut sub: i32 = list[0];
    let mut max_val: i32 = sum;
    let mut max_count: i32 = 1i32;

    for day_group in list.windows(x_n[1]).skip(1) {
        sum -= sub; 
        sub = day_group[0];
        sum += day_group[x_n[1] - 1];
        
        if sum > max_val {
            max_val = sum;
            max_count = 1;
        } else if sum == max_val {
            max_count += 1;
        }
    }

    if max_val == 0 {
        println!("SAD");
    } else {
        println!("{}\n{}", max_val, max_count);
    }
}
