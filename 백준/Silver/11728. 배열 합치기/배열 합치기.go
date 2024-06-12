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

func sToI(s string) int {
	toInt, _ := strconv.Atoi(s)
	return toInt
}

func main() {
	defer writer.Flush()
	readLine()
	A := strings.Fields(readLine())
	B := strings.Fields(readLine())
	result := make([]int, len(A)+len(B))
	for i, data := range append(A, B...) {
		result[i] = sToI(data)
	}
	sort.Ints(result)
	for _, data := range result {
		fmt.Fprint(writer, data, " ")
	}
}
