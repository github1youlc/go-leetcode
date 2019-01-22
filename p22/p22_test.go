/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午10:47
 */

package p22

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	result := generateParenthesis(4)
	for _, r := range result {
		t.Log(r)
	}

}
