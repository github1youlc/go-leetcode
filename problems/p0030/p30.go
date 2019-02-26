/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-31 下午4:20
 */

package p30

func findSubstring(s string, words []string) []int {
	var result []int
	if len(words) == 0 {
		return result
	}

	posIndex := make(map[int][]int)
	length := len(s)

	for i, _ := range s {
		for wi, w := range words {
			if i+len(w) <= length && s[i:i+len(w)] == w {
				posIndex[i] = append(posIndex[i], wi)
			}
		}
	}
	for i, _ := range s {
		found := make([]bool, len(words))
		if isValid(i, words, posIndex, found, 0) {
			result = append(result, i)
		}
	}

	return result
}

func isValid(pos int, words []string, posIndex map[int][]int, found []bool, foundCnt int) bool {
	if len(found) == foundCnt {
		return true
	}
	for _, cand := range posIndex[pos] {
		if !found[cand] {
			found[cand] = true
			if isValid(pos+len(words[cand]), words, posIndex, found, foundCnt+1) {
				return true
			}
			found[cand] = false
		}
	}
	return false
}
