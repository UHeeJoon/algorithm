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
		input   string
		letters [26]int
	)
	fmt.Fscanln(reader, &input)

	input = strings.ToUpper(input)

	for i := 0; i < len(input); i++ {
		letters[input[i]-65]++
	}

	var maxVal = -1
	var maxKey rune
	for key, val := range letters {
		if val > maxVal {
			maxVal = val
			maxKey = rune(key + 65)
		} else if val == maxVal {
			maxKey = '?'
		}
	}

	fmt.Println(string(maxKey))
}
