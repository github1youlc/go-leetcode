package solve

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	t.Log(minSubArrayLen(7, []int{2,3,1,2,4,3}))
	t.Log(minSubArrayLen(15, []int{2,3,1,2,4,3}))
	t.Log(minSubArrayLen(16, []int{2,3,1,2,4,3}))
}
