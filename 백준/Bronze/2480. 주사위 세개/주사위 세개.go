package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b, c, result int
	fmt.Fscan(reader, &a, &b, &c)
	if a == b && b == c {
		result = 10_000 + a*1000
	} else if a == b || b == c {
		result = 1_000 + b*100
	} else if a == c {
		result = 1_000 + a*100
	} else {
		result = max(a, b, c) * 100
	}

	fmt.Fprintln(writer, result)
}
