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
	dx          = [4]int{0, 1, 0, -1}
	dy          = [4]int{1, 0, -1, 0}
	t           int
	row, col, k int
	field       = make([][]bool, row)
)

func search_dfs(y, x int) {
	field[y][x] = false
	for i := 0; i < 4; i++ {
		ny := dy[i] + y
		nx := dx[i] + x
		if ny >= row || ny < 0 || nx >= col || nx < 0 {
			continue
		}
		if field[ny][nx] {
			search_dfs(ny, nx)
		}
	}
}

func main() {
	defer writer.Flush()
	var (
		x, y int
	)
	fmt.Fscanln(reader, &t)
	for tc := 0; tc < t; tc++ {
		fmt.Fscanln(reader, &col, &row, &k)
		var count = 0
		field = make([][]bool, row)
		for i := 0; i < row; i++ {
			field[i] = make([]bool, col)
		}
		for i := 0; i < k; i++ {
			fmt.Fscanln(reader, &x, &y)
			field[y][x] = true
		}
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				if field[i][j] {
					search_dfs(i, j)
					count++
				}
			}
		}
		fmt.Fprintln(writer, count)
	}
}
