package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func get_gcd(a, b int) int {
	if b != 0 {
		return get_gcd(b, a%b)
	}
	return a
}

func solve(a, b int) {
	var gcd = get_gcd(a, b)
	var lcm = a * b / gcd
	fmt.Fprintln(writer, gcd)
	fmt.Fprintln(writer, lcm)
}

func main() {
	defer writer.Flush()
	var num1, num2 int
	fmt.Fscan(reader, &num1, &num2)
	solve(num1, num2)
}
