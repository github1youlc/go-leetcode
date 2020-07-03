package p68

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	ret := make([]string, 0)
	for i := 0; i < n; {
		accLength := len(words[i])
		j := i+1
		for ; j < n && accLength + 1 + len(words[j]) <= maxWidth; j++ {
			accLength += 1 + len(words[j])
		}

		ret = append(ret, adjust(words[i:j], accLength-j+i+1, maxWidth, j==n))
		i=j
	}

	return ret
}

func adjust(words []string, totalLength int,  maxWidth int, lastLine bool) string {
	var  builder strings.Builder
	n := len(words)
	if lastLine {
		builder.WriteString(words[0])
		for i := 1; i  <n ;i ++ {
			builder.WriteByte(' ')
			builder.WriteString(words[i])
		}
		appendEmpty(&builder, maxWidth - totalLength - (n-1))
		return builder.String()
	}


	 if n == 1 {
	 	builder.WriteString(words[0])
		appendEmpty(&builder, maxWidth - totalLength)
		return builder.String()
	 }

	 remain := maxWidth - totalLength

	 cnt := remain / (n - 1)
	 more := remain % (n - 1)

	for i := 0; i < more; i ++ {
		builder.WriteString(words[i])
		appendEmpty(&builder, cnt + 1)
	}

	for i := more; i < n-1; i++ {
		builder.WriteString(words[i])
		appendEmpty(&builder, cnt)
	}

	builder.WriteString(words[n-1])

	return builder.String()
}

func appendEmpty(builder *strings.Builder, n int) {
	for i := 0; i < n ;i ++ {
		builder.WriteByte(' ')
	}
}