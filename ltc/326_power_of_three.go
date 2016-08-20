package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(num int) bool {
	if num <= 0 {
		return false
	}
	return num == int(math.Pow(3, float64(int(math.Log10(float64(num))/math.Log10(3)))))
}

func main() {
	test_case := []int{}
	for i := 0; i < 10; i++ {
		num := int(math.Pow(3, float64(i)))
		test_case = append(test_case, num)
		test_case = append(test_case, num+1)
	}
	for _, num := range test_case {
		fmt.Printf("%v\n", isPowerOfThree(num))
	}
}
