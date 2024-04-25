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

func main() {
	defer writer.Flush()
	var (
		primes = make([]bool, 1010)
		n      = 0
		count  = 0
	)
	primes[1] = true
	for i := 2; i*i <= 1000; i++ {
		if primes[i] {
			continue
		}
		for j := i + i; j <= 1000; j += i {
			primes[j] = true
		}
	}
	fmt.Fscanln(reader, &n)
	input, _ := reader.ReadString('\n')
	nums := strings.Fields(input)
	for _, num := range nums {
		data, _ := strconv.Atoi(num)
		if !primes[data] {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
