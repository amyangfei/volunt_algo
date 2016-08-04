package main

import (
	"container/heap"
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Data struct {
	value int
	x, y  int
}

type DataHeap []*Data

func (h DataHeap) Len() int           { return len(h) }
func (h DataHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h DataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DataHeap) Push(x interface{}) {
	*h = append(*h, x.(*Data))
}

func (h *DataHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	h := &DataHeap{}
	heap.Init(h)

	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	for i := 0; i < len(nums1); i++ {
		heap.Push(h, &Data{nums1[i] + nums2[0], i, 0})
	}

	for len(result) < Min(k, len(nums1)*len(nums2)) {
		d := heap.Pop(h).(*Data)
		result = append(result, []int{nums1[d.x], nums2[d.y]})
		if d.y+1 < len(nums2) {
			heap.Push(h, &Data{nums1[d.x] + nums2[d.y+1], d.x, d.y + 1})
		}
	}

	return result
}
