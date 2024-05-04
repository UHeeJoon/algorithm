package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

var (
	dx      = [4]int{0, 1, 0, -1}
	dy      = [4]int{1, 0, -1, 0}
	RGB     = map[byte]byte{'R': 1, 'G': 2, 'B': 4}
	n       int
	visited [][]bool
	str     []string
)

func visited_init() {
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
}

func dfs(b byte, x, y int) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || nx >= n || ny < 0 || ny >= n || visited[nx][ny] {
			continue
		}
		if RGB[str[nx][ny]]&b == 0 {
			continue
		}
		dfs(b, nx, ny)
	}
}

func search(rg bool) int {
	count := 0
	var b byte
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				continue
			}
			if rg && (str[i][j] == 'R' || str[i][j] == 'G') {
				b = RGB['R'] | RGB['G']
			} else {
				b = RGB[str[i][j]]
			}
			dfs(b, i, j)
			count++
		}
	}
	return count
}

func main() {
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	str = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str[i])
	}
	visited_init()
	fmt.Fprint(writer, search(false), " ")
	visited_init()
	fmt.Fprintln(writer, search(true))
}
