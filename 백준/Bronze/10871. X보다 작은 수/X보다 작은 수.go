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
	var n, x, a int
	fmt.Fscan(reader, &n, &x)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		if a < x {
			fmt.Fprint(writer, a, " ")
		}
	}

}
