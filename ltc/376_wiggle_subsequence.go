package main

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	positive, negative := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			positive = negative + 1
		} else if nums[i] < nums[i-1] {
			negative = positive + 1
		}
	}
	return Max(positive, negative)
}
