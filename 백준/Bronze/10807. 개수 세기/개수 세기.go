package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, x int
	fmt.Fscanln(reader, &n)
	line, _ := reader.ReadString('\n')
	fmt.Fscanln(reader, &x)
	inputs := strings.Fields(line)
	count := 0
	for _, i := range inputs {
		a, _ := strconv.Atoi(i)
		if a == x {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
