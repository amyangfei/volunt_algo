package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, x := range nums {
			if x+i <= target {
				dp[i+x] += dp[i]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}
