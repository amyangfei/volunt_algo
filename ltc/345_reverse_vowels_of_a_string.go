package main

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		for i < j && !isVowel(runes[i]) {
			i++
		}
		for i < j && !isVowel(runes[j]) {
			j--
		}
		if i < j {
			runes[i], runes[j] = runes[j], runes[i]
		}
		i++
		j--
	}
	return string(runes)
}
