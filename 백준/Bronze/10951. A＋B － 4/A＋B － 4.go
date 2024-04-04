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
	var a, b int
	for true {
		val, _ := fmt.Fscan(reader, &a, &b)
		if val != 2 {
			break
		}
		fmt.Fprintln(writer, a+b)
	}

}
