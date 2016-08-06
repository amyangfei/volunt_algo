package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num] = 1
	}
	for _, num := range nums2 {
		if val, ok := m[num]; ok && val == 1 {
			m[num] = 2
		}
	}
	result := []int{}
	for k, v := range m {
		if v == 2 {
			result = append(result, k)
		}
	}
	return result
}
