package p35

func searchInsert(nums []int, target int) int {
	s := 0
	n := len(nums)
	e := n - 1

	var m int
	var mid int
	for s <= e {
		m = s + (e - s) / 2
		mid = nums[m]

		if mid == target {
			return m
		} else if target < mid {
			e = m -1
		} else {
			s = m + 1
		}
	}

	return s
}