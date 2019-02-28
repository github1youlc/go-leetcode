package solve

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	t.Log(rangeBitwiseAnd(5, 7))
	t.Log(rangeBitwiseAnd(7, 10))
	t.Log(rangeBitwiseAnd(0, 1))
}
