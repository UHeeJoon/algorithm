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

func main() {
	defer writer.Flush()
	var str, result string
	fmt.Fscan(reader, &str)

	switch str {
	case "A+":
		result = "4.3"
	case "A0":
		result = "4.0"
	case "A-":
		result = "3.7"
	case "B+":
		result = "3.3"
	case "B0":
		result = "3.0"
	case "B-":
		result = "2.7"
	case "C+":
		result = "2.3"
	case "C0":
		result = "2.0"
	case "C-":
		result = "1.7"
	case "D+":
		result = "1.3"
	case "D0":
		result = "1.0"
	case "D-":
		result = "0.7"
	case "F":
		result = "0.0"
	}
	fmt.Fprintln(writer, result)
}
