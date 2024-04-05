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
	list := [31]bool{false}
	var idx int
	for i := 0; i < len(list)-3; i++ {
		fmt.Fscan(reader, &idx)
		list[idx] = true
	}
	for idx, data := range list {
		if idx != 0 && !data {
			fmt.Fprintln(writer, idx)
		}
	}
}
