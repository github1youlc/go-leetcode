package solve

func minSubArrayLen(s int, nums []int) int {
	//sum := make([]int, len(nums))
	//for l := 0; l< len(nums); l ++ {
	//	for i := 0; i < len(nums) -l; i++ {
	//		sum[i] += nums[i+l]
	//		if sum[i] >= s {
	//			return l + 1
	//		}
	//	}
	//}
	//
	//return 0

	var total, left int
	result := len(nums) + 1
	for right, n := range (nums) {
		total += n
		for total >= s {
			result = intMin(result, right - left + 1)
			total -= nums[left]
			left += 1
		}

	}

	if result <= len(nums) {
		return result
	} else {
		return 0
	}
}

func intMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

