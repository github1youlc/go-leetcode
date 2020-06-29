package p51

import "strings"

func solveNQueens(n int) [][]string {
	mark := [4][]bool{}
	mark[0] = make([]bool, n)
	mark[1] = make([]bool, n)
	mark[2] = make([]bool, 2 * n -1)
	mark[3] = make([]bool, 2 * n -1)

	buf := make([]bool, n * n)
	bss := solve(mark, n, n, 0, buf)

	ret := make([][]string, 0)
	for _, bs := range bss {
		ret = append(ret, boolArray2String(bs, n))
	}
	return ret
}

func solve(mark [4][]bool, left int,  n int, c int, buf []bool) [][]bool {
	if left == 0 {
		s := make([]bool, len(buf))
		copy(s, buf)
		return [][]bool{s}
	}
	if c == n * n {
		return nil
	}

	i := c / n
	j := c % n

	if c % n == 0 && j > 0 && !mark[0][j-1] {
		return nil
	}

	ret := make([][]bool, 0)

	pos1 := i + j
	pos2 := n - 1 -i +j

	if !(mark[0][i] || mark[1][j] || mark[2][pos1] || mark[3][pos2]) {
		mark[0][i] = true
		mark[1][j] = true
		mark[2][pos1] = true
		mark[3][pos2] = true
		buf[c] = true
		ret = append(ret, solve(mark, left -1, n, c+1, buf)...)
		mark[0][i] = false
		mark[1][j] = false
		mark[2][pos1] = false
		mark[3][pos2] = false
		buf[c] = false
	}

	ret = append(ret, solve(mark, left, n, c+1, buf)...)

	return ret
}

func boolArray2String(bs []bool, n int) []string {
	ret := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var buf strings.Builder
		for j := i *n; j < i*n +n; j++ {
			if bs[j] {
				buf.WriteByte('Q')
			} else {
				buf.WriteByte('.')
			}
		}
		ret = append(ret, buf.String())
	}

	return ret
}


