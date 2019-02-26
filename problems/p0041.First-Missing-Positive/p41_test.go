/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午12:23
 */

package p41

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	t.Log(firstMissingPositive([]int{1, 1, 3, 4}))
	t.Log(firstMissingPositive([]int{1, 2, 0}))
	t.Log(firstMissingPositive([]int{3, 4, -1, 1}))
	t.Log(firstMissingPositive([]int{3, 4}))
}
