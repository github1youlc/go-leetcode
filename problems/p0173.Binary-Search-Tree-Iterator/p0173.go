package p0173_Binary_Search_Tree_Iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
	}

	iter.pushStack(root)

	return iter
}

func (this *BSTIterator) pushStack(node *TreeNode) {
	cur := node
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	stackLen := len(this.stack)
	node := this.stack[stackLen - 1]
	this.stack = this.stack[0:stackLen-1]
	if node.Right != nil {
		this.pushStack(node.Right)
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}
