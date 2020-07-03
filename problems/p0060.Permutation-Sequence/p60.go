package p60

import (
	"strings"
)

func getPermutation(n int, k int) string {
	var b strings.Builder
	mark := make([]bool, n)
	return doNext(k - 1, np(n), n, n, mark, &b)
}

func doNext(k int, np int, left int, n int, mark []bool, b *strings.Builder) string {
	if left == 0 {
		return b.String()
	}
	np = np / left
	count := k/np  + 1
	for i := 0; i < n; i++ {
		if !mark[i] {
			count--
		}

		if count == 0 {
			mark[i] = true
			b.WriteByte('1' + byte(i))
			return doNext(k%np, np, left-1, n, mark, b)
		}
	}

	return ""
}

func np(n int) int {
	m := 1

	for i := 1; i <= n; i++ {
		m = m *i
	}

	return m
}