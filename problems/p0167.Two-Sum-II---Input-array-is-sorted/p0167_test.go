package p167

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 18))
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{2, 7, 11, 15}, 17))
	t.Log(twoSum([]int{2, 7, 11, 15}, 15))
}
