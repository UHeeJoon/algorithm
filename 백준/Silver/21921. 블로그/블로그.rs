use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut reader = stdin.lock().lines();

    // 첫 줄 읽기 및 x_n 파싱
    let x_n: Vec<usize> = reader.next().unwrap().unwrap()
        .split_whitespace()
        .flat_map(str::parse::<usize>)
        .collect();

    // 두 번째 줄 읽기 및 list 파싱
    let list: Vec<i32> = reader.next().unwrap().unwrap()
        .split_whitespace()
        .flat_map(str::parse::<i32>)
        .collect();

    let k = x_n[1]; // 슬라이딩 윈도우 크기

    let mut sum: i32 = list[0..k].iter().sum(); // 초기 윈도우 합
    let mut max_val: i32 = sum;
    let mut max_count: i32 = 1;

    // 슬라이딩 윈도우로 최대값 계산
    for i in k..list.len() {
        sum += list[i] - list[i - k]; // 새로운 값 추가, 오래된 값 제거
        if sum > max_val {
            max_val = sum;
            max_count = 1;
        } else if sum == max_val {
            max_count += 1;
        }
    }

    // 결과 출력
    if max_val == 0 {
        println!("SAD");
    } else {
        println!("{}\n{}", max_val, max_count);
    }
}
