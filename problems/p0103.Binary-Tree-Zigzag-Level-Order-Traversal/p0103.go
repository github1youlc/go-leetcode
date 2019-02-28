package solve

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	var levelNodes []*TreeNode
	levelNodes = append(levelNodes, root)
	var reverse = false
	for len(levelNodes) != 0 {
		var nextLevelNodes []*TreeNode
		length := len(levelNodes)
		values := make([]int, length)
		for i, node := range levelNodes {
			if reverse {
				values[length-1-i] = node.Val
			} else {
				values[i] = node.Val
			}

			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		result = append(result, values)
		levelNodes = nextLevelNodes
		reverse = !reverse
	}

	return result
}
