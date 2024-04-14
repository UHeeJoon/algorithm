package main

import (
	"bufio"
	"fmt"

	"os"
)

var writer *bufio.Writer
var dp []int
var num3 []int // 숫자 3개 담아서 정렬할 슬라이스

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines) // 스캐너 옵션 변경
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	X := 0
	fmt.Fscan(reader, &X)

	dp = make([]int, X+1) // 카운트 담을 슬라이스
	dp[1] = 0
	num3 = make([]int, 3) // 숫자 3개 담아서 제일 최소값 찾기 위해 선언

	if X == 1 {
		fmt.Println(dp[1])
		return
	} else if 1 < X && X <= 3 {
		fmt.Println(1)
		return
	} else {
		dp[2] = 1
		dp[3] = 1

		fmt.Println(DynamicProgramming(X))
	}
}

func DynamicProgramming(X int) int {
	if dp[X] != 0 {
		return dp[X]
	}

	if X%6 == 0 {
		dp[X] = minInt(DynamicProgramming(X/3), DynamicProgramming(X/2)) + 1
	} else if X%3 == 0 {
		dp[X] = minInt(DynamicProgramming(X/3), DynamicProgramming(X-1)) + 1
	} else if X%2 == 0 {
		dp[X] = minInt(DynamicProgramming(X/2), DynamicProgramming(X-1)) + 1
	} else {
		dp[X] = DynamicProgramming(X-1) + 1
	}

	return dp[X]
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}