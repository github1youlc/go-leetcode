package p0169_Majority_Element

func majorityElement(nums []int) int {
	num := nums[0]
	count := 0
	for _, n := range nums {
		if n == num {
			count++
		} else {
			count--
		}

		if count < 0 {
			num = n
			count = 1
		}
	}

	return num
}
