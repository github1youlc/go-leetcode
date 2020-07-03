package p66

func plusOne(digits []int) []int {
	n := len(digits)
	ret := make([]int, n+1)

	carry := 1
	digit := 0
	i := n - 1
	for ; i >= 0; i-- {
		if carry == 0 {
			break
		}
		digit = digits[i]
		ret[i+1] = (digit + carry) % 10
		carry = (digit + carry) / 10
	}

	if carry  == 1 {
		ret[0] = 1
		return ret
	} else {
		for j := 0; j <= i; j++ {
			ret[j+1] = digits[j]
		}
		return ret[1:]
	}
}