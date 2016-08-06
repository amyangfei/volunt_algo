package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n >= 10 {
		n = 10
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = 9
	}
	for i := 2; i <= n; i++ {
		for j := 9; j >= 9-i+2; j-- {
			dp[i] *= j
		}
	}
	result := 0
	for i := 0; i <= n; i++ {
		result += dp[i]
	}
	return result
}

func main() {
	for i := 0; i < 12; i++ {
		fmt.Printf("%d\n", countNumbersWithUniqueDigits(i))
	}
}
