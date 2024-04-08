package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var (
		n, sum int
		input  string
	)
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &input)
	for _, data := range []rune(input) {
		sum += int(data - 48)
	}
	fmt.Fprint(writer, sum)
}
