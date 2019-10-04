package p0152_Maximum_Product_SubArray

import (
	"testing"
)

func Test_maxProductNoZero(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	t.Log(a[0:0])

	t.Log(a[3:5])
}

func Test_maxProduct(t *testing.T) {
	//t.Log(maxProduct([]int{0, -1, 0}))

	t.Log(maxProduct([]int{2 ,3,-2,4, 0, 1, 3, 5, 0, 0, -1, 3}))
	t.Log(maxProduct([]int{-2,0,-1}))
}
