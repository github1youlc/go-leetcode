package p167

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l := 0
	r := n - 1

	for l < r {
		cur := numbers[l] + numbers[r]
		if cur == target {
			return []int{l+1, r+1}
		} else if cur < target {
			l++
		} else {
			r--
		}
	}

	return nil
}