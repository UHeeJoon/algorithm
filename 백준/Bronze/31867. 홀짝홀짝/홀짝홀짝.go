package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return input
}

func input() string {
	_ = readLine()                       // n isn't used
	return strings.TrimSpace(readLine()) // return k
}

func countNumbers(num int, o, e *int) {
	if (num & 1) == 1 {
		*o++ // increase odd number
		return
	}
	*e++ // increase even number
}

func countOddAndEven(str string) (int, int) {
    odd, even := 0, 0 // variables
	for _, num := range str {
		countNumbers(int(num), &odd, &even)
	}
	return odd, even
}

func isOddGreaterThanEven(odd, even int) int {
	if even < odd {
		return 1
	}
	if odd < even {
		return 0
	}
	return -1
}

func main() {
	defer writer.Flush()
	odd, even := countOddAndEven(input())
	result := isOddGreaterThanEven(odd, even)
	fmt.Fprintln(writer, result)
}
