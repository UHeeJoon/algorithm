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

func stringToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func solve() {
	sum := 0
	tmp := 0
	for i := 0; i < 5; i++ {
		tmp = stringToInt(readLine())
		if tmp < 40 {
			tmp = 40
		}
		sum += tmp
	}
	fmt.Fprintln(writer, sum/5)
}

func main() {
	defer writer.Flush()
	solve()
}
