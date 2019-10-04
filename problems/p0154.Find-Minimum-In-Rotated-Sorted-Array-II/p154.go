package p0153_Find_Minimum_In_Rotated_Sorted_Array

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return minInt(nums[0], nums[1])
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
		rv := nums[right]

		if lv == rv {
			return minInt(findMin(nums[left: mid+1]), findMin(nums[mid:right+1]))
		}

		if lv <= mv {
			left = mid
		} else {
			right = mid
		}

		if right-left <= 1 {
			return nums[right]
		}
	}

	return -1
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}