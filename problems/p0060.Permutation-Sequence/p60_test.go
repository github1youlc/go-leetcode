package p60

import "testing"

func Test_getPermutation(t *testing.T) {
	t.Log(getPermutation(3, 1))
	t.Log(getPermutation(3, 2))
	t.Log(getPermutation(3, 3))
	t.Log(getPermutation(3, 4))
	t.Log(getPermutation(3, 5))
	t.Log(getPermutation(3, 6))
	t.Log(getPermutation(3, 7))
	t.Log(getPermutation(4, 9))
}