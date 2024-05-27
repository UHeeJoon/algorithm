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

func stringToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func getNK() (int, int) {
	nk := strings.Fields(readLine())
	return stringToInt(nk[0]), stringToInt(nk[1])
}

func getCoins(n int) []int {
	coin := make([]int, n)
	for i := 0; i < n; i++ {
		coin[i] = stringToInt(readLine())
	}
	return coin
}

func query(coin, dp []int, n, k int) {
	for i := 0; i < n; i++ {
		for j := coin[i]; j <= k; j++ {
			dp[j] += dp[j-coin[i]]
		}
	}
}

func solve() {
	n, k := getNK()
	dp := make([]int, k+1)
	dp[0] = 1
	coin := getCoins(n)
	query(coin, dp, n, k)
	fmt.Fprintln(writer, dp[k])
}

func main() {
	defer writer.Flush()
	solve()
}
