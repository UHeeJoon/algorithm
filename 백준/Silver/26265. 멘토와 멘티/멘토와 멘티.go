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

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func splitLine(line string) (string, string) {
	lineArr := strings.Fields(line)
	return lineArr[0], lineArr[1]
}

type Point struct {
	x, y string
}

type Points []Point

func (p Points) Len() int      { return len(p) }
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		return p[i].y > p[j].y
	}
	return p[i].x < p[j].x
}
func main() {
	defer writer.Flush()
	n, _ := strconv.Atoi(readLine())

	points := make(Points, n)
	for i := 0; i < n; i++ {
		x, y := splitLine(readLine())
		points[i] = Point{x, y}
	}

	sort.Sort(points)
	for _, point := range points {
		fmt.Fprintln(writer, point.x, point.y)
	}
}
