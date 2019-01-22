/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-31 下午3:41
 */

package p42

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	maxIndex, max := maxHeightIndex(height)

	if max < height[0] && max <height[length-1]  {
		return calculateTrap(height)
	} else {
		return trap(height[0:maxIndex+1]) + trap(height[maxIndex:])
	}
}

func maxHeightIndex(height []int) (int, int) {
	max := height[1]
	maxIndex := 1
	for i := 1; i < len(height) - 1; i++ {
		if height[i] > max {
			maxIndex = i
			max = height[i]
		}
	}

	return maxIndex, max
}

func calculateTrap(height []int) int {
	length := len(height)
	min := intMin(height[0], height[length-1])
	result := 0
	for i := 1; i < length-1; i++ {
		result += min - height[i]
	}

	return result
}

func intMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
