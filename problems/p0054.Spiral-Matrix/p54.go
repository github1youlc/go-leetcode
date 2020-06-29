package p54

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	return doSpiralOrder(matrix, 0, 0, m-1, n-1)
}


func doSpiralOrder(matrix [][]int, x int, y int, m int, n int) []int {
	if x > m || y > n {
		return nil
	}

	if x == m {
		return matrix[x][y:n+1]
	}

	ret := make([]int, 0)
	if y == n {
		for i := x; i <= m; i++ {
			ret = append(ret, matrix[i][y])
		}

		return ret
	}

	i := x
	j := y
	for ;j < n; j++ {
		ret = append(ret, matrix[i][j])
	}

	for ;i < m; i++ {
		ret = append(ret, matrix[i][j])
	}

	for ;j > x; j-- {
		ret = append(ret, matrix[i][j])
	}

	for ;i > y; i-- {
		ret = append(ret, matrix[i][j])
	}

	return append(ret, doSpiralOrder(matrix, x+1, y+1, m-1, n-1)...)
}