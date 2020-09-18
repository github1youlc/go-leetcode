package solve

type Node struct {
	Next []*Edge
}

type Edge struct {
	Prev *Node
	Next *Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {

}