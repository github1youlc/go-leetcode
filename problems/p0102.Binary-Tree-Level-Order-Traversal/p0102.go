package  solve


//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	var  levelNodes []*TreeNode
	levelNodes = append(levelNodes, root)
	for len(levelNodes) != 0 {
		var values []int
		var nextLevelNodes []*TreeNode
		for _, node := range levelNodes {
			values = append(values, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		result = append(result, values)
		levelNodes = nextLevelNodes
	}

	return result
}