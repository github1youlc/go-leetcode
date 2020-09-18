package p221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])

	dp := make([][][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]bool, n)

	}
}