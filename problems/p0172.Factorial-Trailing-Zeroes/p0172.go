package p0172_Factorial_Trailing_Zeroes

func trailingZeroes(n int) int {
	result := 0

	for n > 0 {
		result += n / 5
		n = n / 5
	}

	return result
}