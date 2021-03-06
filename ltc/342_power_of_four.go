package main

import "fmt"

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 != 0
}

func main() {
	fmt.Printf("%v\n", isPowerOfFour(16))
	fmt.Printf("%v\n", isPowerOfFour(166))
}
