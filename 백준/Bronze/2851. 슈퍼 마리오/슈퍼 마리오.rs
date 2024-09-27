use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut sum: i32 = 0i32;
    let mut flag: bool = true;
    (0..10).for_each(|_| {
        let mut num: String = String::new();
        reader.read_line(&mut num).unwrap();
        let cur_num: i32 = num.trim().parse::<i32>().unwrap();
        if flag && sum < 100 {
            let cur_value_remaining_up_to_100: i32 = 100 - sum;
            if cur_value_remaining_up_to_100 >= cur_value_remaining_up_to_100.abs_diff(cur_num) as i32 {
                sum += cur_num;
            } else {
                flag = false;
            }
        }
    });
    println!("{sum}");
}
