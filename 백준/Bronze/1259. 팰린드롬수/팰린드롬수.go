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
	for {
		var (
			str    string
			result = "yes"
		)
		fmt.Fscan(reader, &str)
		if str == "0" {
			break
		}
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-i-1] {
				result = "no"
				break
			}
		}
		fmt.Fprintln(writer, result)
	}
}
