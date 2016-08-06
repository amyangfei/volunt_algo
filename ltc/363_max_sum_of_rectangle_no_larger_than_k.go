package main

/*
https://discuss.leetcode.com/topic/48875/accepted-c-codes-with-explanation-and-references

find the max sum rectangle in 2D array:
	https://www.youtube.com/watch?v=yCQN096CwWM
find the max sum no more than K in an array:
	https://www.quora.com/Given-an-array-of-integers-A-and-an-integer-k-find-a-subarray-that-contains-the-largest-sum-subject-to-a-constraint-that-the-sum-is-less-than-k
*/

import (
	"math"
	"sort"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func InsertIgnore(s *[]int, x int) {
	i := sort.Search(len(*s), func(i int) bool { return (*s)[i] >= x })
	if i == len(*s) {
		*s = append(*s, x)
	} else if (*s)[i] != x {
		*s = append(*s, 0)
		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = x
	}
}

func LowerBound(s []int, x int) (bool, int) {
	i := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if i < len(s) {
		return true, s[i]
	} else {
		return false, 0
	}
}

func NewSlice(n int, value int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = value
	}
	return s
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}

	row, col := len(matrix), len(matrix[0])
	res := math.MinInt64

	for l := 0; l < col; l++ {
		sums := NewSlice(row, 0)
		for r := l; r < col; r++ {
			for i := 0; i < row; i++ {
				sums[i] += matrix[i][r]
			}

			accus := []int{0}
			curSum, curMax := 0, math.MinInt64
			for _, sum := range sums {
				curSum += sum
				if found, lb := LowerBound(accus, curSum-k); found {
					curMax = Max(curMax, curSum-lb)
				}
				InsertIgnore(&accus, curSum)
				res = Max(res, curMax)
			}
		}
	}
	return res
}
