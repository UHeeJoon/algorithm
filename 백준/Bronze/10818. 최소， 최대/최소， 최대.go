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
	var n int
	_max := -1000001
	_min := 1000001
	fmt.Fscanln(reader, &n)
	input, _ := reader.ReadString('\n')
	data_list := strings.Fields(input)
	for _, data := range data_list {
		el, _ := strconv.Atoi(data)
		if el > _max {
			_max = el
		}
		if el < _min {
			_min = el
		}
	}
	fmt.Fprint(writer, _min, _max)
}
