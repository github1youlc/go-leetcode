package solve

/*
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	rootVal := postorder[len(postorder) - 1]

	root := &TreeNode{
		Val:rootVal,
	}

	i := 0

	for ; inorder[i] != rootVal; i++ {

	}

	root.Left = buildTree(inorder[0:i], postorder[0:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}
