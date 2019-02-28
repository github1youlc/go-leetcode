package solve

/*
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}

	}

	root := &TreeNode{
		Val: preorder[0],
	}

	i := 0
	for ; inorder[i] != preorder[0]; i++ {

	}

	root.Left = buildTree(preorder[1:i+1], inorder[0: i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
