package p64

import "testing"

func Test_minPathSum(t *testing.T) {
	grid := [][]int {
		{1 ,3 ,1},
		{1, 5, 1},
		{4, 2, 1},
	}

	t.Log(minPathSum(grid))
}