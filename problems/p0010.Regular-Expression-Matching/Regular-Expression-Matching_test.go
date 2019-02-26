/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-6 下午12:14
 */

package solve

import "testing"

func Test_isMatch(t *testing.T) {
	if isMatch("aa", "a") {
		t.Error()
	}

	if !isMatch("aa", "a*") {
		t.Error()
	}

	if !isMatch("aaab", "c*a*b") {
		t.Error()
	}

	if isMatch("mississippi", "mis*is*p*.") {
		t.Error()
	}
}
