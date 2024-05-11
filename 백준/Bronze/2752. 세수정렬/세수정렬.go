package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	input, _ := reader.ReadString('\n')
	data := strings.Fields(input)
	sort.Slice(data, func(i, j int) bool {
		i2, _ := strconv.Atoi(data[i])
		j2, _ := strconv.Atoi(data[j])
		return i2 < j2
	})
	for _, v := range data {
		fmt.Fprint(writer, v, " ")
	}
}
