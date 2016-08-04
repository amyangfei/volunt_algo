package main

import "fmt"
import "math"

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = 0
		}
	}

	for i := 2; i < n+1; i++ {
		for j := i - 1; j > 0; j-- {
			globalMin := math.MaxInt64
			for k := j + 1; k < i; k++ {
				localMax := k + Max(dp[j][k-1], dp[k+1][i])
				globalMin = Min(localMax, globalMin)
			}
			if j+1 == i {
				dp[j][i] = j
			} else {
				dp[j][i] = globalMin
			}
		}
	}

	return dp[1][n]
}

func main() {
	fmt.Println(getMoneyAmount(7))
}
