package main

import (
	"bufio"
	"fmt"
	"os"
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

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func splitLine(line string) (string, string) {
	lineArr := strings.Fields(line)
	return lineArr[0], lineArr[1]
}

func initFriendsArray(n int) []string {
	friends := make([]string, n)
	return friends
}

func inputFriendsInfo(friends []string, n int) {
	for i := 0; i < n; i++ {
		friends[i] = readLine()
	}
}

func initCountsArray(n int) [][]int {
	counts := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		counts[i] = make([]int, n+1)
	}
	return counts
}

func floyd(friends []string, fCountsArray [][]int, n int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if friends[i][j] == 'Y' || (friends[i][k] == 'Y' && friends[k][j] == 'Y') {
					fCountsArray[i][j] = 1
				}
			}
		}
	}
}

func sum(row []int) int {
	result := 0
	for i := 0; i < len(row); i++ {
		result += row[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getResult(fCountsArray [][]int) int {
	result := 0
	for _, row := range fCountsArray {
		result = max(result, sum(row))
	}
	return result
}

func main() {
	defer writer.Flush()
	n := strToInt(readLine())
	friends := initFriendsArray(n)

	inputFriendsInfo(friends, n)

	fCountsArray := initCountsArray(n)

	floyd(friends, fCountsArray, n)

	result := getResult(fCountsArray)
	fmt.Fprintln(writer, result)
}
