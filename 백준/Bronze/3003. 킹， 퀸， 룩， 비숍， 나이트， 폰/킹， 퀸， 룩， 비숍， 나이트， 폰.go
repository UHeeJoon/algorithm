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
	chess := []int{1, 1, 2, 2, 2, 8}
	input, _ := reader.ReadString('\n')
	list := strings.Fields(input)
	for idx, v := range list {
		data, _ := strconv.Atoi(v)
		fmt.Fprint(writer, chess[idx]-data, " ")
	}
}
