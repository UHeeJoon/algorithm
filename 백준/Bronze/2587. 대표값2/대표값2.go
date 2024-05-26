package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	arr := make([]int, 5)
	avg := 0
	for i := 0; i < 5; i++ {
		arr[i] = stringToInt(readLine())
		avg += arr[i]
	}
	sort.Ints(arr)
	fmt.Fprintln(writer, avg/5)
	fmt.Fprintln(writer, arr[2])
}

func main() {
	defer writer.Flush()
	solve()
}
