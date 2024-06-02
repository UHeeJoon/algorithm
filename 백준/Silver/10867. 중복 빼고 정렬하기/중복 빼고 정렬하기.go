package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func sToI(s string) int {
	toInt, _ := strconv.Atoi(s)
	return toInt
}

func main() {
	defer writer.Flush()
	var arr = make([]int, 2001)
	readLine()
	input := strings.Fields(readLine())
	for _, data := range input {
		arr[sToI(data)+1000] = 1
	}
	for idx, data := range arr {
		if data == 1 {
			fmt.Fprint(writer, idx-1000, " ")
		}
	}
}
