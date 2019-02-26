/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-3 下午6:37
 */

package solve

import "testing"

func Test_convert(t *testing.T) {
	r := convert("PAYPALISHIRING", 3)
	if r != "PAHNAPLSIIGYIR" {
		t.Error(r)
	}

	r = convert("PAYPALISHIRING", 4)
	if r != "PINALSIGYAHRPI" {
		t.Error(r)
	}

	t.Log(convert("A", 1))
}
