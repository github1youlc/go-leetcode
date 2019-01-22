/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-21 上午12:58
 */

package p27

import "testing"

func Test_removeElement(t *testing.T) {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	l := removeElement(a, 2)

	for i := 0; i < l; i++ {
		t.Log(a[i])
	}
}
