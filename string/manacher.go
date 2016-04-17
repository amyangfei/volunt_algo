/*
Manacher algorithm
*/
package string

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Manacher(s string) string {
	// prepare string
	rs := []rune(s)
	prep := []rune("$")
	for _, ch := range rs {
		prep = append(prep, []rune{rune('#'), rune(ch)}...)
	}
	prep = append(prep, rune('#'))
	prep = append(prep, rune('$'))

	// calculate p[i]
	mx, id := 0, 0
	p := make([]int, len(prep))
	for i := 1; i < len(prep); i++ {
		if mx > i {
			p[i] = Min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}

		for i-p[i] > 0 && i+p[i] < len(prep) && prep[i+p[i]] == prep[i-p[i]] {
			p[i]++
		}
		if p[i]+i > mx {
			mx = p[i] + i
			id = i
		}
	}

	// find max p[i]
	mid, maxp := 0, 0
	for i := 0; i < len(prep); i++ {
		if p[i] > maxp {
			maxp = p[i]
			mid = i
		}
	}

	// get result
	result := []rune{}
	if prep[mid] == rune('#') {
		for i := 1; i < maxp; i += 2 {
			result = append([]rune{prep[i+mid]}, result...)
			result = append(result, prep[i+mid])
		}
	} else {
		result = []rune{prep[mid]}
		for i := 2; i < maxp; i += 2 {
			result = append([]rune{prep[i+mid]}, result...)
			result = append(result, prep[i+mid])
		}
	}
	return string(result)
}
