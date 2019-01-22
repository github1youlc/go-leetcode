/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-26 下午9:44
 */

package p18

import (
	"testing"
)

func Test_fourSum(t *testing.T) {
	result := fourSum([]int{-3,-2,-1,0,0,1,2,3}, 0)

	for _, r := range result {
		t.Log(r)
	}
}

func testMap(t *testing.T)  {

}