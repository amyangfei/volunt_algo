package main

/* O(klogn)。使用最小堆，进行k-1此pop操作，去除k-1小的元素 */

import (
	"container/heap"
	"fmt"
)

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

func kthSmallest(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])

	h := &DataHeap{}
	heap.Init(h)
	for i := 0; i < col; i++ {
		heap.Push(h, &Data{matrix[0][i], 0, i})
	}

	for j := 0; j < k-1; j++ {
		curr := heap.Pop(h).(*Data)
		x, y := curr.x, curr.y
		if x+1 < row {
			heap.Push(h, &Data{matrix[x+1][y], x + 1, y})
		}
	}

	return heap.Pop(h).(*Data).value
}

func main() {
	var matrix [][]int
	matrix = append(matrix, []int{1, 5, 9})
	matrix = append(matrix, []int{10, 11, 13})
	matrix = append(matrix, []int{12, 13, 15})
	fmt.Println(kthSmallest(matrix, 8))
	// expected 13
}
