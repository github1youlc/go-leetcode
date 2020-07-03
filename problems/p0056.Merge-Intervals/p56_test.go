package p56

import (
	"testing"
)

func Test_merge(t *testing.T) {
	{
		intervals := [][]int{
			{1, 3}, {2, 6}, {8, 10}, {15, 18},
		}

		ret := merge(intervals)

		t.Log(ret)
	}

	{
		intervals := [][]int{
			{1, 4}, {2, 3},
		}

		ret := merge(intervals)

		t.Log(ret)
	}
}
