package string

func Reverse(s []rune, start, end int) {
	// Reverse rune from start to end-1
	for i, j := start, end-1; i < start+(end-start)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseWord(word string) string {
	s := []rune(word)
	i, j := 0, -1
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			if j+1 < i-1 {
				Reverse(s, j+1, i)
			}
			j = i
		}
	}
	if j+1 < i-1 {
		Reverse(s, j+1, i)
	}
	Reverse(s, 0, len(s))
	return string(s)
}
