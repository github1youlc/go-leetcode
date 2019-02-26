/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午12:58
 */

package p27

func removeElement(nums []int, val int) int {
	length := len(nums)

	validLength := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[validLength] = nums[i]
			validLength++
		}
	}

	return validLength
}
