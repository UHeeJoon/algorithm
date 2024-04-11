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
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var h, w, n int
		fmt.Fscan(reader, &h, &w, &n)
		var room, floor int
		floor = n % h
		room = n/h + 1
		if floor == 0 {
			floor = h
			room -= 1
		}
		fmt.Fprintln(writer, floor*100+room)
	}
}
