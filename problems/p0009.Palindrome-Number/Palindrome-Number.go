/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-3 下午9:08
 */

package solve

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	return x == reverse(x)
}

func reverse(x int) int {
	r := 0
	for x > 0 {
		r = r * 10 + x % 10
		x = x / 10
	}

	return r
}

func checkPalindrome(x int) bool {
	length := 0
	powLeft := 1
	cur := x
	for cur > 0 {
		length++
		powLeft *= 10
		cur = cur / 10
	}

	powLeft /= 10
	var left, right int
	for i := 0; i < length/2; i++ {
		right = x % 10
		left = x / powLeft % 10
		if left != right {
			return false
		}

		x = x %powLeft / 10
		powLeft = powLeft / 100
	}

	return true
}
