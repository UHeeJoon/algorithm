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

// 입력 받고 공백 제거
func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// 문자열 정수로
func stringToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

// 원소 입력 받고 누적합 배열 만들기
func makeDPFromInput() []int {
	_ = readLine() // n은 필요 없음
    // 배열 입력
	arr := strings.Fields(readLine())
    // 누적 합 배열
	dp := make([]int, len(arr)+1)
	for i, data := range arr {
        // dp[cur] = dp[prev] + arr[cur]
		dp[i+1] = dp[i] + stringToInt(data)
	}
	return dp
}

// query입력 받아서 int로 변환 후 return
func getQuery() (int, int) {
	query := strings.Fields(readLine())
	return stringToInt(query[0]), stringToInt(query[1])
}

// query 실행
func query(dp []int) {
    // testcase
	ts := stringToInt(readLine())
	for i := 0; i < ts; i++ {
		q1, q2 := getQuery()
		fmt.Fprintln(writer, dp[q2]-dp[q1-1])
	}
}

func main() {
	defer writer.Flush()
	query(makeDPFromInput())
}
