package p53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}


	m := make([]int, len(nums))
	m[0] = nums[0]
	max := m[0]

	for i := 1; i < len(nums); i++ {
		if m[i-1] > 0 {
			m[i] = m[i-1] + nums[i]
		} else {
			m[i] = nums[i]
		}

		if m[i] > max {
			max = m[i]
		}
	}
	return max
}

