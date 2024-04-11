package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_self_num(self_num []bool, num int) {
	tmp := num
	for num > 0 {
		tmp += num % 10
		num /= 10
	}
	self_num[tmp] = true
}
func main() {
	// reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	self_num := make([]bool, 10037)
	for i := 1; i < 10000+1; i++ {
		get_self_num(self_num, i)
		if self_num[i] {
			continue
		}
		fmt.Fprintln(writer, i)
	}
}
