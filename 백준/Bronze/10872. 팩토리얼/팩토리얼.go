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
	var n int
	facto := 1
	fmt.Fscan(reader, &n)
	for i := 2; i <= n; i++ {
		facto *= i
	}
	fmt.Fprintln(writer, facto)

}
