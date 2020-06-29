package p35

import "testing"

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5,6}
	t.Log(searchInsert(nums, 0))
	t.Log(searchInsert(nums, 1))
	t.Log(searchInsert(nums, 2))
	t.Log(searchInsert(nums, 3))
	t.Log(searchInsert(nums, 4))
	t.Log(searchInsert(nums, 5))
	t.Log(searchInsert(nums, 6))
	t.Log(searchInsert(nums, 7))
}