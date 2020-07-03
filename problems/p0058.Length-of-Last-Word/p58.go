package p58


func lengthOfLastWord(s string) int {
	n := len(s)

	pos := n -1

	for ; pos >= 0; pos -- {
		if s[pos] != ' ' {
			break
		}
	}

	length := 0
	for ; pos >= 0; pos -- {
		if s[pos] != ' ' {
			length++
		} else {
			break
		}
	}

	return length
}