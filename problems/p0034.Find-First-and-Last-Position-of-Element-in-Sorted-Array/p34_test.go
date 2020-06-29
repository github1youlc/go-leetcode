package p34

import (
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5,7,7,8,8,10}

	t.Log(searchRange(nums, 8))
	t.Log(searchRange(nums, 5))
	t.Log(searchRange(nums, 7))
	t.Log(searchRange(nums, 10))
	t.Log(searchRange(nums, 9))
}