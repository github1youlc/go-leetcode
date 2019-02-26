/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-3 下午6:37
 */

package solve

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var builder strings.Builder
	period := 2 * (numRows - 1)
	length := len(s)
	var right int
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < length; j += period {
				builder.WriteByte(s[j])
			}
		} else {
			for j := i; j < length; j += period {
				builder.WriteByte(s[j])
				right = period+j-2*i
				if right < length {
					builder.WriteByte(s[right])
				}
			}
		}
	}
	return builder.String()
}
