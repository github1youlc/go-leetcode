/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-26 下午9:44
 */

package p18

import "sort"

type IntSortList []int

func (s IntSortList) Len() int {
	return len(s)
}

func (s IntSortList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSortList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int

	sort.Sort(IntSortList(nums))

	length := len(nums)

	var i, j, k, m, s int
	var kVal, mVal int
	for i = 0; i < length-3; i++ {
		for j = i + 1; j < length-2; j ++ {
			k = j + 1
			m = length - 1
			for k < m {
				s = nums[i] + nums[j] + nums[k] + nums[m]
				if s == target {
					result = insertUnique(result, []int{nums[i], nums[j], nums[k], nums[m]}, 4)

					kVal = nums[k]
					k++
					for nums[k] == kVal && k < m {
						k++
					}

					mVal = nums[m]
					m--
					for nums[m] == mVal && k < m {
						m--
					}
				} else if s < target {
					k++
				} else {
					m--
				}
			}
		}
	}
	return result
}

func insertUnique(result [][]int, item []int, length int) [][]int {

	var start, end int
	start = 0
	end = len(result) - 1

	for r := 0; r < length; r ++ {
		for start <= end && result[start][r] < item[r] {
			start++
		}
		for start <= end && result[end][r] > item[r] {
			end --
		}
	}

	if start != end {
		tmp := result[start:]
		result = append(result[0:start], item)
		result = append(result, tmp...)
	}

	return result
}
