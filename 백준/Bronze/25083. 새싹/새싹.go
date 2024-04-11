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
	str :=
		"         ,r'\"7\n" +
			"r`-_   ,'  ,/\n" +
			" \\. \". L_r'\n" +
			"   `~\\/\n" +
			"      |\n" +
			"      |\n"
	fmt.Fprintln(writer, str)
}
