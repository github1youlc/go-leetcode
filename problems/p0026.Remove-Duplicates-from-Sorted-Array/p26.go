/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午9:23
 */

package p26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prev := nums[0]
	length := 1
	for _, d := range nums {
		if d != prev {
			nums[length] = d
			prev = d
			length = length + 1
		}
	}

	return length
}
