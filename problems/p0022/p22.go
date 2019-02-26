/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-14 下午10:47
 */

package p22

func generateParenthesis(n int) []string {
	return gen(n * 2, 0)
}

func gen(n int, lCnt int) []string {
	if n == 0 {
		return []string{""}
	}

	var result []string
	if lCnt == 0 {
		p := "("
		rem := gen(n, lCnt+1)
		for _, l := range rem {
			result = append(result, p+l)
		}
		return result

	} else if lCnt >= 0 && n-lCnt > lCnt {
		p := "("
		rem := gen(n, lCnt+1)
		for _, l := range rem {
			result = append(result, p+l)
		}
		p = ")"
		rem = gen(n-2, lCnt-1)
		for _, l := range rem {
			result = append(result, p+l)
		}
		return result
	} else {
		p := ")"
		rem := gen(n-2, lCnt-1)
		for _, l := range rem {
			result = append(result, p+l)
		}
		return result
	}
}
