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
		input string
		i     int
	)
	fmt.Fscanln(reader, &input)
	fmt.Fscanln(reader, &i)
	fmt.Fprintln(writer, string([]rune(input)[i-1]))
}
