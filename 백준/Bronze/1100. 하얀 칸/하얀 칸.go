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
	var count int
	for i := 0; i < 8; i++ {
		input, _ := reader.ReadString('\n')
		datas := strings.Split(input, "")
		for j := 0; j < len(datas); j++ {
			if datas[j] == "F" {
				if ((i&1) == 0 && (j&1) == 0) || ((i&1) == 1 && (j&1) == 1) {
					count++
				}
			}
		}
	}
	fmt.Fprintln(writer, count)
}
