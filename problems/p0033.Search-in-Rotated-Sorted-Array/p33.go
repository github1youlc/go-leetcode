package p33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	start := nums[0]
	end := nums[n-1]
	if target > end && target < start {
		return -1
	}

	m := n / 2
	mid := nums[m]

	if mid == target {
		return m
	} else if (mid >= start && (start <= target && target < mid)) ||
		(mid <= end && !(mid < target && target <= end)) {
		return search(nums[0:m], target)
	} else {
		ret := search(nums[m+1:n], target)
		if ret == -1 {
			return -1
		} else {
			return m + 1 + ret
		}
	}
}
