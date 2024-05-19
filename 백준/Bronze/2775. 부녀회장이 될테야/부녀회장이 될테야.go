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

func main() {
	defer writer.Flush()
	var a, b, c int
	fmt.Fscanln(reader, &a)
	var A = make([][]int, 15)
	for i := 0; i < 15; i++ {
		A[i] = make([]int, 15)
		A[0][i] = i + 1

	}
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			d := 0
			for k := 0; k <= j; k++ {
				d += A[i][k]
			}
			A[i+1][j] = d
		}
	}
	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &b)
		fmt.Fscanln(reader, &c)
		fmt.Fprintln(writer, A[b][c-1])
	}
}
