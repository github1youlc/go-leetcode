/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午12:23
 */

package p41

func firstMissingPositive(nums []int) int {
	length := len(nums)
	var i int
	for i = range nums {
		for nums[i] > 0 && nums[i] <= length && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i] - 1]
		} 
	}
	
	for i = 0;  i < length; i++ {
		if nums[i] != i + 1 {
			return i+1
		}
	}
	
	return length + 1
}