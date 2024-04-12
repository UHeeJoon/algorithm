package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	var list = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &list[i])
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	for _, data := range list {
		fmt.Fprintln(writer, data)
	}
}
