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
	var (
		list = make([][]string, 9)
	)
	for i := 0; i < 9; i++ {
		input, _ := reader.ReadString('\n')
		list[i] = strings.Fields(input)
	}
	_max := -99999999
	locate_x, locate_y := 0, 0
	for i, el := range list {
		for j, data := range el {
			int_data, _ := strconv.Atoi(data)
			if int_data > _max {
				_max = int_data
				locate_y, locate_x = i+1, j+1
			}
		}
	}
	fmt.Fprintln(writer, _max)
	fmt.Fprintln(writer, locate_y, locate_x)
}
