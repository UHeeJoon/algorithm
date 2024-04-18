package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func max(a, b int) int {
	if a >b {
		return a
	}
	return b
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m  int
	i:=0 
	j:=0 
	sum:=0
	result := -999999999
	fmt.Fscanln(reader, &n, &m)
	input, _:=reader.ReadString('\n')
	arr := strings.Fields(input)
	for j < n{
		if j - i < m {
			el, _ := strconv.Atoi(arr[j])
			sum += el
			j++
		} 
		if j- i >= m {
			result = max(result, sum)
			el, _ := strconv.Atoi(arr[i])
			sum -= el
			i++
		}
	}
	fmt.Fprintln(writer, result)
}