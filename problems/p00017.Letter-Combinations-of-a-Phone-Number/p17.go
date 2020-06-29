package p17

var (
	phoneMap = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	chars := convertDigits(digits)
	bs := traverse(chars, 0, len(chars), nil)
	ret := make([]string, 0, len(bs))
	for _, b := range bs {
		ret = append(ret, string(b))
	}

	return ret
}

func convertDigits(digits string) [][]byte {
	result := make([][]byte, 0, len(digits))
	for _, b := range digits {
		if c, ok := phoneMap[uint8(b)]; ok {
			result = append(result, c)
		}
	}

	return result
}

func traverse(chars [][]byte, pos int, len int, prefix [][]byte) [][]byte {
	if pos == len {
		return prefix
	}

	return traverse(chars, pos+1, len, multiply(prefix, chars[pos]))
}

func multiply(prefix [][]byte, cur []byte) [][]byte {
	result := make([][]byte, 0, len(prefix)*len(cur))

	if len(prefix) == 0 {
		for _, c := range cur {
			result = append(result, []byte{c})
		}
	} else {
		for _, p := range prefix {
			for _, c := range cur {
				t := append([]byte{}, p...)
				result = append(result, append(t, c))
			}
		}
	}

	return result
}
