package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "|\\_/|")
	fmt.Fprintln(writer, "|q p|   /}")
	fmt.Fprintln(writer, "( 0 )\"\"\"\\")
	fmt.Fprintln(writer, "|\"^\"`    |")
	fmt.Fprintln(writer, "||_/=\\\\__|")
}
