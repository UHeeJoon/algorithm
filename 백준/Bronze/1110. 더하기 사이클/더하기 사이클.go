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

func cycle(a int) int {
	var sum = 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	return sum % 10
}

func calcul(a int) int {
	return a%10*10 + cycle(a)
}

func main() {
	defer writer.Flush()
	var n, count int = 0, 1
	fmt.Fscan(reader, &n)
	var tmp = calcul(n)
	for n != tmp {
		tmp = calcul(tmp)
		count++
	}
	fmt.Fprintln(writer, count)
}
