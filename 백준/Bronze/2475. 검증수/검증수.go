package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input, _ := reader.ReadString('\n')
	nums := strings.Fields(input)
	sum := 0.0
	for _, num := range nums {
		data, _ := strconv.Atoi(num)
		pow := math.Pow(float64(data), 2)
		sum += pow
	}
	fmt.Fprintln(writer, int64(sum)%10)
}
