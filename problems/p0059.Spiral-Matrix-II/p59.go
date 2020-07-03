package p59

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	doFill(1, n, 0, ret)

	return ret
}

func doFill(num int, left int, start int, ret [][]int) {
	if left == 0 {
		return
	}

	if left == 1 {
		ret[start][start] = num
		return
	}

	i := start
	j := start

	for ; j < start+ left - 1; j++ {
		fill(ret, i, j, &num)
	}

	for ; i < start + left - 1; i++ {
		fill(ret, i ,j ,&num)
	}

	for ; j > start; j-- {
		fill(ret, i, j, &num)
	}

	for ; i > start; i-- {
		fill(ret, i, j, &num)
	}

	doFill(num, left - 2, start+ 1, ret)
}

func fill(ret [][]int ,i int, j int , num *int) {
	ret[i][j] = *num
	*num = *num + 1
}
