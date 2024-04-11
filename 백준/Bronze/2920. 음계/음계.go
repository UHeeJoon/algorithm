package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	isAscending, isDescending := true, true
	input, _ := reader.ReadString('\n')
	datas := strings.Fields(input)
	for i := 1; i < len(datas); i++ {
		if datas[i] > datas[i-1] {
			isDescending = false
		} else if datas[i] < datas[i-1] {
			isAscending = false
		}
	}
	if isDescending {
		fmt.Fprintln(writer, "descending")
	} else if isAscending {
		fmt.Fprintln(writer, "ascending")
	} else {
		fmt.Fprintln(writer, "mixed")
	}
}
