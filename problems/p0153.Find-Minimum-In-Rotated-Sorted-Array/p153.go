package p0153_Find_Minimum_In_Rotated_Sorted_Array

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if nums[0] < nums[n-1] {
		return nums[0]
	}

	left := 0
	right := n - 1



	for true {
		mid := (left + right) / 2
		lv := nums[left]
		mv := nums[mid]

		if lv > mv {
			right = mid
		} else {
			left = mid
		}

		if right - left <= 1 {
			return nums[right]
		}
	}

	return 0
}