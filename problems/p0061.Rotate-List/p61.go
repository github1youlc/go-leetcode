package p61

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    n := 0
    for cur := head; cur != nil; cur = cur.Next {
        n++
    }

    if n == 0 {
        return head
    }

    k = k % n
    if k == 0 {
        return head
    }

    c1 := head
    for i := 0; i < k; i++ {
        c1 = c1.Next
    }

    c2 := head
    for c1.Next != nil {
        c1 = c1.Next
        c2 = c2.Next
    }

    newHead := c2.Next
    c1.Next = head
    c2.Next = nil

    return newHead
}