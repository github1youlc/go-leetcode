package solve

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}

	diff := n - m
	topDiff := 1
	for topDiff <= diff {
		topDiff <<= 1
	}

	return rangeBitwiseAnd(m / topDiff, n / topDiff) * topDiff
}
