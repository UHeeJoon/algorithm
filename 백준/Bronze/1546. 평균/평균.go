package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMax(list []string) float64 {
	data := 0.0
	for _, v := range list {
		_v, _ := strconv.ParseFloat(v, 64)
		if _v > data {
			data = _v
		}
	}
	return data
}

func sum(list []string) float64 {
	var res float64
	max := getMax(list)
	for _, v := range list {
		data, _ := strconv.ParseFloat(v, 64)
		res += (data / max) * 100
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	n, _ := reader.ReadString('\n')
	len, _ := strconv.ParseFloat(strings.TrimSpace(n), 64)
	input, _ := reader.ReadString('\n')
	fmt.Fprintf(writer, "%0.3f\n", sum(strings.Fields(input))/len)
}
