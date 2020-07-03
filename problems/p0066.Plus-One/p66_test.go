package p66

import (
	"testing"
)

func Test_plusOne(t *testing.T) {
	t.Log(plusOne(nil))
	t.Log(plusOne([]int{9, 9, 9}))
	t.Log(plusOne([]int{7, 9}))
	t.Log(plusOne([]int{7, 5, 3, 4}))
}