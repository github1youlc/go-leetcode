package solve

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	head := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}


	printList(removeElements(head, 6))
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
