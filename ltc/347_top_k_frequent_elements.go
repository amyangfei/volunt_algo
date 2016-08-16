package main

import (
	"container/heap"
	"fmt"
)

type Data struct {
	value, freq int
}

type DataHeap []*Data

func (h DataHeap) Len() int           { return len(h) }
func (h DataHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h DataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DataHeap) Push(x interface{}) {
	*h = append(*h, x.(*Data))
}

func (h *DataHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	h := &DataHeap{}
	heap.Init(h)

	for kk, v := range m {
		if len(*h) < k {
			heap.Push(h, &Data{kk, v})
		} else if (*h)[0].freq < v {
			heap.Pop(h)
			heap.Push(h, &Data{kk, v})
		}
	}

	result := make([]int, k)
	i := 0
	for len(*h) > 0 {
		d := heap.Pop(h).(*Data)
		result[i] = d.value
		i++
	}

	return result
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Printf("%v\n", topKFrequent(arr, k))
}
