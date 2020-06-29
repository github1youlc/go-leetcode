/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-31 下午4:20
 */

package p30

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return nil
	}

	width := len(words[0])
	number := len(words)
	window := width * number

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	length := len(s)
	var ret []int

	for i := 0; i < width; i++ {
		start := i
		for start+window <= length {
			findAll := true
			countCopy := copyWordCount(wordCount)
			for j := start; j < start+window; j += width {
				word := s[j : j+width]
				v, ok := countCopy[word]
				if !ok {
					findAll = false
					start = j
					break
				}

				if v <= 0 {
					findAll = false
					break
				} else {
					countCopy[word]--
				}
			}

			if findAll {
				ret = append(ret, start)
			}

			start += width
		}
	}

	return ret
}

func copyWordCount(wordCount map[string]int) map[string]int {
	ret := make(map[string]int)

	for key, value := range wordCount {
		ret[key] = value
	}
	return ret
}
