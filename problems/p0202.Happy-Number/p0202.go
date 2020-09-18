package solve

func isHappy(n int) bool {
	traversed := make(map[int]struct{})

	cur := n
	var next int
	for true {
		if cur == 1 {
			return true
		}

		_, found := traversed[cur]
		if found {
			return false
		}

		traversed[cur] = struct{}{}

		next = 0
		for cur != 0 {
			next += cur % 10 * (cur % 10)
			cur = cur / 10
		}

		cur = next
	}

	return false
}
