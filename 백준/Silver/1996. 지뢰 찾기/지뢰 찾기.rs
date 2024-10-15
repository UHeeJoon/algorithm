use std::io::{self, BufReader, BufRead};


fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();
    let len: usize = input.trim().parse::<usize>().unwrap();
    let mut map: Vec<Vec<char>> = vec![];
    let mut ans: Vec<Vec<char>> = vec![vec!['\0';len];len];
    (0..len).for_each(|_| {
        input.clear();
        reader.read_line(&mut input).unwrap();
        map.push(input.trim().chars().collect::<Vec<char>>());
    });

    const dy:[i32; 8] =[0,0,1,1,1,-1,-1,-1]; 
    const dx:[i32; 8] =[1,-1,0,1,-1,0,1,-1]; 

    (0..len).for_each(|i| {
        (0..len).for_each(|j| {

            let mut counter: i32 = 0i32;
            if map[i][j] == '.' {
                for k in 0..8 {
                    let iy = i as i32 + dy[k];
                    let jx = j as i32 + dx[k];
                    if iy < 0 || iy >= len as i32 || jx < 0 || jx >= len as i32 || map[iy as usize][jx as usize] == '.' {
                        continue;
                    }
                    counter += map[iy as usize][jx as usize].to_digit(10).unwrap() as i32;
                }
                ans[i][j] = char::from_digit(counter as u32, 10).unwrap_or('M');
            } else {
                ans[i][j] = '*';
            }

        });
    });
    
    ans.iter().for_each(|str| 
        println!("{}", str.iter().collect::<String>())
    );
}
