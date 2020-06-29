package p37

type pos struct {
	i int
	j int
}

func solveSudoku(board [][]byte)  {
	toFill := make([]*pos, 0)
	status := [3][9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				toFill = append(toFill, &pos{i: i, j:j})
			} else {
				n := board[i][j] - '1'
				status[0][i][n] = true
				status[1][j][n] = true
				status[2][calcCell(i, j)][n] = true

			}
		}
	}

	_ = dfs(board, toFill, 0, status)
}

func calcCell(i, j int) int{
	return i / 3 * 3 + j / 3
}

func dfs(board [][]byte, toFill []*pos, filledCount int, status [3][9][9]bool) bool {
	if filledCount == len(toFill) {
		return true
	}

	p := toFill[filledCount]
	for n := 0; n < 9; n++ {
		cell := calcCell(p.i, p.j)
		if !status[0][p.i][n] && !status[1][p.j][n] && !status[2][cell][n] {
			status[0][p.i][n] = true
			status[1][p.j][n] = true
			status[2][cell][n] = true
			board[p.i][p.j] = byte(n) + '1'
 			if dfs(board, toFill, filledCount+1, status) {
 				return true
			}
			status[0][p.i][n] = false
			status[1][p.j][n] = false
			status[2][cell][n] = false
		}
	}

	return false
}

