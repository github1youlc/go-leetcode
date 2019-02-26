package solve

func isSubsequence(s string, t string) bool {
	pos := 0
	tLen := len(t)
	matchCount := 0
	for i := 0; i < len(s); i ++ {
		sc := s[i]
		for pos < tLen {
			if t[pos] == sc {
				matchCount++
				pos ++
				break
			} else {
				pos ++
			}
		}
	}
	return matchCount == len(s)
}

