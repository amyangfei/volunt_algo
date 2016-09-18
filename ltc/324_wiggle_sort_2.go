package main

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	n := len(nums)
	mid := (n + 1) / 2
	last := n
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			mid -= 1
			temp[i] = nums[mid]
		} else {
			last -= 1
			temp[i] = nums[last]
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = temp[i]
	}
}
