use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut n: String = String::new();
    let mut input: String = String::new();
    reader.read_line(&mut n).unwrap();
    reader.read_line(&mut input).unwrap();

    let n: usize = n.trim().parse::<usize>().unwrap();
    let mut graph: Vec<Vec<i32>> = vec![vec![0; n + 1]; n + 1];
    let mut visited: Vec<bool> = vec![false; n + 1];

    (0..input.trim().parse().unwrap()).for_each(|_| {
        let mut cp = String::new();
        reader.read_line(&mut cp).unwrap();
        let list: Vec<i32> = cp.split_whitespace().flat_map(str::parse::<i32>).collect::<Vec<i32>>();
        graph[list[0] as usize].push(list[1]);
        graph[list[1] as usize].push(list[0]); 
    });

    visited[1] = true;
    let answer: i32 = search(&graph, &mut visited, 1); 
    println!("{}", answer);
}

fn search(graph: &Vec<Vec<i32>>, visited: &mut Vec<bool>, idx: usize) -> i32 {
    let mut ans: i32 = 0;
    for ele in graph[idx].clone() {
        if ele == 0 {
            continue;
        }
        if !visited[ele as usize] {
            visited[ele as usize] = true;
            ans += 1 + search(graph, visited, ele as usize);
        }
    }
    ans
}