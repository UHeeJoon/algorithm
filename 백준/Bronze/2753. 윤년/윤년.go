package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(num int) int {
	if (num%4 == 0 && num%100 != 0) || (num%400 == 0) {
		return 1
	}
	return 0

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input, _ := reader.ReadString('\n')
	data, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Fprint(writer, check(data))
}
