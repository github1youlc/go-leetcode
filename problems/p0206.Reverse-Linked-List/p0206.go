package solve

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var newCur *ListNode
	cur := head.Next
	head.Next = nil
	for cur != nil {
		newCur = cur.Next
		cur.Next = head
		head = cur
		cur = newCur
	}
	return head
}