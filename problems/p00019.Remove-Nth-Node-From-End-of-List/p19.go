package p19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}

	cur2 := head
	for i := 0; i < n-1 && cur2 != nil; i++ {
		cur2 = cur2.Next
	}

	if cur2 == nil {
		return head
	}

	if cur2.Next == nil {
		return head.Next
	}

	cur2 = cur2.Next
	cur := head
	for cur2.Next != nil {
		cur = cur.Next
		cur2 = cur2.Next
	}

	cur.Next = cur.Next.Next

	return head
}
