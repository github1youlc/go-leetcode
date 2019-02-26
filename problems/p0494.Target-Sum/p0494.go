/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-25 下午8:41
 */

package p0494_Target_Sum

import "math"

func findTargetSumWays(nums []int, S int) int {
	length := len(nums)
	sum := intSum(nums)
	if (sum + S) % 2 != 0 || sum < S {
		return 0
	}

	//return  dfs(nums, S, 0, length)

	if math.Pow(2, float64(length)) < float64(S) * float64(length) {
		return dfs(nums, S, 0, length)
	} else {
		return subSet(nums, (S + sum) / 2)
	}
}

func dfs(nums[]int, s int, cur int, length int) int {
	if cur == length {
		if s == 0 {
			return 1
		} else {
			return 0
		}
	}

	return dfs(nums, s - nums[cur], cur + 1, length) + dfs(nums, s + nums[cur], cur + 1, length)
}

func intSum(nums[]int) int {
	s := 0

	for _, num := range nums {
		s += num
	}

	return s
}

func subSet(nums []int, target int) int  {
	dp := make([]int, target + 1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}