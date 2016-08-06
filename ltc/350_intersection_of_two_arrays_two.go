package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	result := []int{}
	for _, num := range nums2 {
		if val, ok := m[num]; ok {
			result = append(result, num)
			if val > 1 {
				m[num]--
			} else {
				delete(m, num)
			}
		}
	}
	return result
}
