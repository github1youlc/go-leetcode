/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午1:07
 */

package p28

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}

	if haystackLen < needleLen {
		return -1
	}


	for i := 0; i <= haystackLen - needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
