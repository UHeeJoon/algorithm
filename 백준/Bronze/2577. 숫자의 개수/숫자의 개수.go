package main

import (
	"bufio"
	"fmt"
	"os"
)

func divide_digits(number int) [10]int {
	nums := [10]int{0}
	for number != 0 {
		nums[number%10]++
		number /= 10
	}
	return nums
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)
	for _, data := range divide_digits(a * b * c) {
		fmt.Fprintln(writer, data)
	}
}
