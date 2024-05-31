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
	padoban := [102]int{0, 1, 1, 1, 2, 2, 3, 4, 5, 7, 9}
	for i := 11; i < 101; i++ {
		padoban[i] = padoban[i-2] + padoban[i-3]
	}
	ts := sToI(readLine())
	for ts > 0 {
		n := sToI(readLine())
		fmt.Fprintln(writer, padoban[n])
		ts--
	}
}
