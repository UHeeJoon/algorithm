package main

import (
	"bufio"
	"fmt"
	"os"
)

type Fibonacci struct {
	zero int64
	one  int64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var t, n int
	fmt.Fscan(reader, &t)
	list := make([]Fibonacci, 41)
	list[0] = Fibonacci{1, 0}
	list[1] = Fibonacci{0, 1}
	list[2] = Fibonacci{1, 1}
	list[3] = Fibonacci{1, 2}
	for i := 4; i < 41; i++ {
		list[i].zero = list[i-1].zero + list[i-2].zero
		list[i].one = list[i-1].one + list[i-2].one
	}

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, list[n].zero, list[n].one)
	}
}
