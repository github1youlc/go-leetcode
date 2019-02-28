package solve

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	root := buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
	t.Log(root)
}
