package p61

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {
	{
		l := slice2List([]int{1, 2, 3,4 ,5})

		t.Log(list2Slice(rotateRight(l, 2)))
	}

	{
		{
			l := slice2List([]int{0, 1, 2})

			t.Log(list2Slice(rotateRight(l, 1)))
		}
	}
}

func slice2List(s []int) *ListNode {
	head := &ListNode{}

	cur := head
	for _, v := range s {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return head.Next
}

func list2Slice(head *ListNode) []int {
	ret := make([]int, 0)

	for cur := head; cur != nil; cur= cur.Next {
		ret = append(ret, cur.Val)
	}

	return ret
}