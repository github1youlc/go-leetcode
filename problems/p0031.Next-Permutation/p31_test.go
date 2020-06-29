package p31

import "testing"

func Test_nextPermutation(t *testing.T) {
	nums := []int{2, 3, 1, 3, 3}

	nextPermutation(nums)

	t.Log(nums)
}