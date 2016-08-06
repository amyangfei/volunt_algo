package main

/*
z = m * x + n * y（m与n为整数，可以为0或者为负）
令x = p * gcd, y = q * gcd，可知p与q互质。
则z = (m * p + n * q) * gcd

可以证明一定存在m, n，使得m * p + n * q = 1（p与q互质的性质，参见：裴蜀定理）
由此可知z一定是gcd的整数倍
*/

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func canMeasureWater(x int, y int, z int) bool {
	return (x+y == z) || ((x+y > z) && z%gcd(x, y) == 0)
}
