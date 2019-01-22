/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午1:19
 */

package p5

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	maxLen := 1
	result := s[0:1]

	var l, begin, end int
	for i := 0; i < len(s); i++ {
		if i < len(s) - 1 && s[i] == s[i+1] {
			l := 0
			begin = i - 1
			end = 2 * i + 1 - begin
			for begin >= 0 && end < len(s) && s[begin] == s[end] {
				l += 1
				begin = begin - 1
				end = 2 * i + 1 - begin
			}

			if 2 + 2*l > maxLen {
				result = s[i-l: i+l+2]
				maxLen = 2 + 2 * l
			}
		}

		l = 0
		begin = i - 1
		end = 2 * i - begin
		for begin >= 0 && end < len(s) && s[begin] == s[end] {
			l += 1
			begin = begin - 1
			end = 2 * i - begin
		}

		if 1 + 2 * l > maxLen {
			result = s[i-l: i+l+1]
			maxLen = 1 + 2 * l
		}
	}

	return result
}
