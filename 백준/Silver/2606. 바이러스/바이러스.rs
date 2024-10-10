use std::io::{self, BufRead, BufReader};

fn main() {
    let reader = BufReader::new(io::stdin().lock());
    let mut lines = reader.lines();

    let n: usize = lines.next().unwrap().unwrap().trim().parse().unwrap();
    let input_count: usize = lines.next().unwrap().unwrap().trim().parse().unwrap();

    let mut graph = vec![vec![]; n + 1];
    let mut visited = vec![false; n + 1];

    for _ in 0..input_count {
        if let Some(Ok(line)) = lines.next() {
            let list: Vec<usize> = line.split_whitespace()
                                       .map(|s| s.parse().unwrap())
                                       .collect();
            if list.len() == 2 {
                let (u, v) = (list[0], list[1]);
                graph[u].push(v);
                graph[v].push(u);
            }
        }
    }

    visited[1] = true;
    let answer = dfs(&graph, &mut visited, 1);
    println!("{}", answer);
}

fn dfs(graph: &[Vec<usize>], visited: &mut [bool], idx: usize) -> i32 {
    graph[idx]
        .iter()
        .filter_map(|&neighbor| {
            if !visited[neighbor] {
                visited[neighbor] = true;
                Some(1 + dfs(graph, visited, neighbor))
            } else {
                None
            }
        })
        .sum()
}