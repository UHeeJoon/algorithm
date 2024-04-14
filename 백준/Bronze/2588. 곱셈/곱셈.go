package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var A, B int
	fmt.Fscan(reader, &A, &B)
	tmp := B
	for tmp > 0 {
		fmt.Fprintln(writer, A*(tmp%10))
		tmp /= 10
	}
	fmt.Fprintln(writer, A*B)
}
