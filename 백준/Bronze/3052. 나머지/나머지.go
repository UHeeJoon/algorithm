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
	var (
		nums  [42]bool
		input int
		count int
	)
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &input)
		nums[input%42] = true
	}
	for _, data := range nums {
		if data {
			count++
		}
	}
	fmt.Fprint(writer, count)
}
