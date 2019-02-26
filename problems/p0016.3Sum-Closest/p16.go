/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-26 下午9:17
 */

package p16

import "sort"

type IntSortList []int

func (s IntSortList) Len() int {
	return len(s)
}

func (s IntSortList) Less(i, j int) bool  {
	return s[i] < s[j]
}

func (s IntSortList) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Sort(IntSortList(nums))

	var i, j, k, tmp, result int
	result = nums[0] + nums[1] + nums[2]
	for i = 0; i < len(nums) - 2; i++ {
		j = i + 1
		k = len(nums) - 1
		for j < k {
			tmp = nums[i] + nums[j] + nums[k]
			if tmp == target {
				return tmp
			}
			if closer(target, tmp, result) {
				result = tmp
			}
			if tmp > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func closer(target, num1, num2 int) bool {
	return diff(target, num1) < diff(target, num2)
}

func diff(n1, n2 int) int {
	d := n1 - n2
	if d >= 0 {
		return d
	} else {
		return -d
	}
}