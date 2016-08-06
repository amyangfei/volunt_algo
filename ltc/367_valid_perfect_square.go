package main

/*
Fast inverse square root:
https://en.wikipedia.org/wiki/Fast_inverse_square_root
*/

/*
#include <stdint.h>

int32_t FastInvSqrt_C(int32_t x) {
    double y=x;
    int64_t i=0x5fe6eb50c7b537a9;
    i = i-(*(int64_t*)&y)/2;
    y = *(double*)&i;
    y = y * (3 - x * y * y) * 0.5;
    y = y * (3 - x * y * y) * 0.5;
    i = x * y + 1;
    return i - (i * i > x);
}
*/
import "C"
import "fmt"
import "unsafe"

func FastInvSqrt(x float32) float32 {
	xhalf := float32(0.5) * x
	y := x
	i := *(*int32)(unsafe.Pointer(&y))
	i = int32(0x5f3759df) - int32(i>>1)
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (1.5 - (xhalf * y * y))
	return y
}

func isPerfectSquare2(num int) bool {
	start, end := 1, num
	for start <= end {
		mid := (start + end) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	root := int(C.FastInvSqrt_C(C.int32_t(num)))
	return root*root == num
}

func main() {
	var num float64 = 602176
	fmt.Println(isPerfectSquare(int(num)))
}
