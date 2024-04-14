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
	for i := 0; i < 2; i++ {
		fmt.Fprintln(writer, "강한친구 대한육군")
	}

}
