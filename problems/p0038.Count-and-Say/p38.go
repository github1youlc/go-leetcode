package p38

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	prev := countAndSay(n-1)

	cur := prev[0]
	count := 0

	var ret bytes.Buffer
	var ch byte
	for _, c := range prev {
		ch = byte(c)
		if ch  == cur {
			count++
		} else {
			ret.Write([]byte(strconv.Itoa(count)))
			ret.WriteByte(cur)

			cur = ch
			count = 1
		}
	}

	ret.Write([]byte(strconv.Itoa(count)))
	ret.WriteByte(cur)

	return ret.String()
}
