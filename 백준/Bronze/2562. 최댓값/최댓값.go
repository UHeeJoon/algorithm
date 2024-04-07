package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_max, idx := 0, 0
	for i := 0; i < 9; i++ {
		num, _ := reader.ReadString('\n')
		data, _ := strconv.Atoi(strings.Fields(num)[0])
		if _max < data {
			_max = data
			idx = i + 1
		}
	}
	fmt.Fprintln(writer, _max)
	fmt.Fprintln(writer, idx)
}
