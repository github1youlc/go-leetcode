/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-25 下午1:13
 */

package p0347_Top_K_Frequent_Elements

import (
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	t.Log(topKFrequent([]int{1,1,1,2,2,2,23}, 2))
	t.Log(topKFrequent([]int{1}, 1))
}
