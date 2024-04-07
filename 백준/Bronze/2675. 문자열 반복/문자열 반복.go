package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var (
		t, r int
		s    string
	)
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &r, &s)
		for _, data := range strings.Split(s, "") {
			fmt.Fprint(writer, strings.Repeat(data, r))
		}
		fmt.Fprintln(writer)

	}
}
