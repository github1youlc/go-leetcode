package p57

import (
	"testing"
)

func Test_insert(t *testing.T) {
	//{
	//	 intervals := [][]int{
	//		{1, 3}, {6, 9},
	//	}
	//	newInterval := []int{
	//		2, 5,
	//	}
	//
	//	ret := insert(intervals, newInterval)
	//	t.Log(ret)
	//}

	{
		intervals := [][]int{
			{1, 5},
		}
		newInterval := []int{
			2, 3,
		}

		ret := insert(intervals, newInterval)
		t.Log(ret)
	}
}