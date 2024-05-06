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

func stoi(a, b string) (int, int) {
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)
	return num1, num2
}

func main() {
	defer writer.Flush()
	reader.ReadString('\n')
	input_A, _ := reader.ReadString('\n')
	input_B, _ := reader.ReadString('\n')
	A := strings.Fields(input_A)
	B := strings.Fields(input_B)
	sort.Slice(A, func(i, j int) bool {
		num1, num2 := stoi(A[i], A[j])
		return num1 < num2
	})
	sort.Slice(B, func(i, j int) bool {
		num1, num2 := stoi(B[i], B[j])
		return num1 > num2
	})
	sum := 0
	for i := 0; i < len(A); i++ {
		num1, num2 := stoi(A[i], B[i])
		sum += num1 * num2
	}
	fmt.Fprintln(writer, sum)
}
