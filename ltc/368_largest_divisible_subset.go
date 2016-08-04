package main

import "fmt"
import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	dp := make([]int, n)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i], idx[i] = 1, -1
	}

	sort.Ints(nums)
	maxDp, maxIdx := 1, 0
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				idx[i] = j
			}
		}
		if maxDp < dp[i] {
			maxDp, maxIdx = dp[i], i
		}
	}

	result := make([]int, 0)
	for maxIdx != -1 {
		// result = append(result, nums[maxIdx])
		result = append(result[:0], append([]int{nums[maxIdx]}, result[0:]...)...)
		maxIdx = idx[maxIdx]
	}
	return result
}

func main() {
	nums := []int{8, 4, 2, 1}
	fmt.Printf("%v\n", largestDivisibleSubset(nums))
}
