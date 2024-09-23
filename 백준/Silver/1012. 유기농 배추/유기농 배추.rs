use std::io::{self, BufRead, BufReader, Write};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    io::stdout().flush().unwrap();
    let mut t: String = String::new();
    reader.read_line(&mut t).unwrap();

    let mut t: i32 = t.trim().parse().unwrap();
    while { t -= 1; t } >= 0 {
        let mut input: String = String::new();
        reader.read_line(&mut input).unwrap();

        let input: Vec<i32> = input.split_whitespace().map(|x| x.trim().parse::<i32>().unwrap()).collect::<Vec<i32>>();
        let (m, n, mut k) = (input[0], input[1], input[2]);

        let mut arr: Vec<Vec<bool>> = vec![vec![false; m as usize]; n as usize];
        while { k -= 1; k } >= 0 {
            let mut input: String = String::new();
            reader.read_line(&mut input).unwrap();

            let input = input.split_whitespace().map(|x| x.trim().parse::<usize>().unwrap()).collect::<Vec<usize>>();
            arr[input[1]][input[0]] = true;
        }

        let di: [i32; 4] = [0, 1, 0, -1];
        let dj: [i32; 4] = [1, 0, -1, 0];
        let mut count: i32 = 0;

        (0..n).for_each(|i| {
            (0..m).for_each(|j| {
                let i = i as usize;
                let j = j as usize;
                if arr[i][j] {
                    search_dfs(m, n, &mut arr, &di, &dj, i, j);
                    count += 1;
                }
            })
        });
        println!("{}", count);
    }
}

fn search_dfs(m: i32, n: i32, arr: &mut Vec<Vec<bool>>, di: &[i32; 4], dj: &[i32; 4], i: usize, j: usize) {
    arr[i][j] = false; // 현재 위치를 방문 처리
    for idx in 0..4 {
        let ni = i as i32 + di[idx];
        let nj = j as i32 + dj[idx];

        if ni < 0 || ni >= n || nj < 0 || nj >= m {
            continue;
        }

        let ni = ni as usize;
        let nj = nj as usize;

        if arr[ni][nj] {
            search_dfs(m, n, arr, di, dj, ni, nj);
        }
    }
}
