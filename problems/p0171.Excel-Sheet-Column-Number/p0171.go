package p0171_Excel_Sheet_Column_Number

func titleToNumber(s string) int {
	leng := len(s)

	result := 0
	pow := 1
	for i := leng - 1; i >= 0; i-- {
		result += pow * char2Int(s[i])
		pow *= 26
	}

	return result
}

func char2Int(c uint8) int {
	return int((c - 'A') + 1)
}