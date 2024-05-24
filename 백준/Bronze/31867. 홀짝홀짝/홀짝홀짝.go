package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	odd, even, ans := 0, 0, -1
	for _, num := range input {
		if (num & 1) == 1 {
			odd++
		} else {
			even++
		}
	}
	if even < odd {
		ans = 1
	} else if odd < even {
		ans = 0
	}
	fmt.Fprintln(writer, ans)
}
