use std::{collections::VecDeque, io::{self, BufRead, BufReader}};

fn main() {
    let reader = BufReader::new(io::stdin().lock());
    let mut lines = reader.lines();
    
    let (n, m) = {
        let size = lines.next().unwrap().unwrap();
        let mut iter = size.split_ascii_whitespace().map(str::parse::<i32>);
        (iter.next().unwrap().unwrap(), iter.next().unwrap().unwrap())
    };

    let mut list: Vec<Vec<u32>> = lines.take(n as usize).map(|line| {
        line.unwrap().chars().map(|c| c.to_digit(10).unwrap()).collect()
    }).collect();

    let mut que: VecDeque<(i32, i32, i32)> = VecDeque::new();
    que.push_back((0, 0, 1));
    list[0][0] = 0;

    let dy = [0, 1, 0, -1];
    let dx = [1, 0, -1, 0];
    
    while let Some((cur_y, cur_x, steps)) = que.pop_front() {
        if cur_y == n - 1 && cur_x == m - 1 {
            println!("{}", steps);
            return;
        }

        for i in 0..4 {
            let ny = cur_y + dy[i];
            let nx = cur_x + dx[i];

            if ny >= 0 && ny < n && nx >= 0 && nx < m && list[ny as usize][nx as usize] == 1 {
                list[ny as usize][nx as usize] = 0; // 방문 처리
                que.push_back((ny, nx, steps + 1));
            }
        }
    }

}
