/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-31 下午3:03
 */

package p8

const (
	int31bit = 2147483648
)

func myAtoi(str string) int {
	length := len(str)

	startIdx := 0
	for startIdx < length && str[startIdx] == ' ' {
		startIdx++
	}
	//
	//for startIdx < length && str[startIdx] == '0' {
	//	startIdx++
	//}

	str = str[startIdx:]

	if str == "" {
		return 0
	} else if str[0] == '-' {
		return int(-parseNum(str[1:], int31bit))
	} else if str[0] == '+' {
		return int(parseNum(str[1:], int31bit-1))
	} else {
		return int(parseNum(str, int31bit-1))
	}
}

func parseNum(str string, max int64) int64 {
	var num int64 = 0
	var n int64
	for _, c := range str {
		n = int64(c - '0')
		if n < 0 || n > 9 {
			return num
		}
		num = num*10 + n
		if num > max {
			return max
		}
	}

	return num
}
