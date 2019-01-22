/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午9:23
 */

package p26

import "testing"

func Test_removeDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 2, 3, 3, 4, 4}
	length := removeDuplicates(nums)
	for i := 0 ; i < length; i++ {
		t.Log(nums[i])
	}
}
