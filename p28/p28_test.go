/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午1:07
 */

package p28

import "testing"

func Test_strStr(t *testing.T) {
	t.Log(strStr("hello", "ll"))
	t.Log(strStr("aaaaaa", "aab"))
	t.Log(strStr("aa", "aab"))
	t.Log(strStr("aaa", ""))
}
