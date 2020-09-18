package solve

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	father := &ListNode{Next: head}
	prev := father

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			cur = cur.Next
			prev = prev.Next
		}
	}

	return father.Next
}
