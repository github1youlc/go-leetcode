package p19

import (
	"fmt"
	"testing"
)

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

func Test_removeNthFromEnd(t *testing.T) {
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 0)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 1)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 2)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 3)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 4)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 5)))
	}
	{
		head := array2List([]int{1, 2, 3, 4, 5})
		fmt.Println(list2Array(removeNthFromEnd(head, 6)))
	}

}
