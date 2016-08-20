package main

import "fmt"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i < n+1; i++ {
		for j := 1; j < i; j++ {
			temp := dp[i-j] * j
			if temp > dp[i] {
				dp[i] = temp
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Printf("%d\n", integerBreak(4))
}
