package p55


func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	last := n - 1
	for i := n-2; i >= 0; i-- {
		if (i + nums[i]) >= last {
			last = i
		}
	}

	return last == 0
}
