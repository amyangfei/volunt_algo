package main

import (
	"fmt"
	"sort"
)

type Pnt struct {
	x, y int
}

type Pnts []Pnt

func (p Pnts) Len() int {
	return len(p)
}

func (p Pnts) Less(i, j int) bool {
	// Attention here!!!
	// one envelope into another, must be greater than, not no less than
	return p[i].x < p[j].x || (p[i].x == p[j].x && p[i].y > p[j].y)
}

func (p Pnts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	pts := Pnts{}
	for idx, _ := range envelopes {
		pts = append(pts, Pnt{envelopes[idx][0], envelopes[idx][1]})
	}
	sort.Sort(pts)
	lis := []int{}

	for i := 0; i < len(envelopes); i++ {
		pos := sort.Search(len(lis), func(j int) bool { return lis[j] >= pts[i].y })
		if pos >= len(lis) {
			lis = append(lis, pts[i].y)
		} else {
			lis[pos] = pts[i].y
		}
	}

	return len(lis)
}

func main() {
	arr := [][]int{
		// 4
		// []int{4, 5},
		// []int{4, 6},
		// []int{6, 7},
		// []int{2, 3},
		// []int{1, 1},

		// 3
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}
	fmt.Printf("%d\n", maxEnvelopes(arr))
}
