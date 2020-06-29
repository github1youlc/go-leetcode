package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func array2List(a []int) *ListNode {
	p := &ListNode{}

	n := p
	for _, ai := range a {
		n.Next = &ListNode{
			Val: ai,
		}
		n = n.Next
	}

	return p.Next
}

func list2Array(head *ListNode) []int {
	result := make([]int, 0)
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}