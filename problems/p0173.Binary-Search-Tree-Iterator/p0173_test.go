package p0173_Binary_Search_Tree_Iterator

import "testing"

func TestBSTIterator_HasNext(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Val: 3,
		},
		Val: 7,
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	iter := Constructor(root)

	t.Log(iter.Next())
	t.Log(iter.Next())
	t.Log(iter.HasNext())
	t.Log(iter.Next())
	t.Log(iter.HasNext())
	t.Log(iter.Next())
	t.Log(iter.HasNext())
	t.Log(iter.Next())
	t.Log(iter.HasNext())
}