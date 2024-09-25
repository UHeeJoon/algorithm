use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    let mut x_point: Vec<i32> = vec![];
    let mut y_point: Vec<i32> = vec![];
    for _ in 0..3 {
        input.clear();
        reader.read_line(&mut input).unwrap();
        let point: Vec<i32> = input.trim().split_whitespace().map(|x| x.trim().parse::<i32>().unwrap()).collect();

        update_point(&mut x_point, point[0]);
        update_point(&mut y_point, point[1]);
    }
    println!("{} {}", x_point[0], y_point[0]);
}

fn update_point(x_point: &mut Vec<i32>, point: i32) {
    match x_point.iter().position(|&x| x == point) {
        Some(idx) => { x_point.remove(idx); }
        None => { x_point.push(point); }
    }
}
