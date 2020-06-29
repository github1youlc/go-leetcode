package p54

import (
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	{
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		t.Log(spiralOrder(matrix))
	}

	{
		matrix := [][]int{
			{1, 2},
			{5, 6},
			{9, 10},
		}

		t.Log(spiralOrder(matrix))
	}
}