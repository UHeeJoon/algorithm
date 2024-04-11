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
	var (
		_time, _minute int
	)
	fmt.Fscanf(reader, "%d %d", &_time, &_minute)
	if _minute-45 < 0 {
		_time -= 1
	}
	_minute = (_minute + 60 - 45) % 60
	_time = (_time + 24) % 24
	fmt.Fprint(writer, _time, _minute)
}
