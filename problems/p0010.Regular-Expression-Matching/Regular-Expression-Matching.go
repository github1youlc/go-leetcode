/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-6 下午12:14
 */

package solve

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	if sLen == 0 && pLen == 0 {
		return true
	}

	if sLen == 0 && pLen != 0 {
		return false
	}

	if sLen != 0 && pLen == 0 {
		return false
	}

	sF := s[0]
	pF := p[0]

	if pLen == 1 {
		if sLen != 1 {
			return false
		}

		return pF == '.' || pF == sF
	}

	pS := p[1]
	if pS == '*' {
		if isMatch(s, p[2:]) {
			return true
		}
		
		if sF == pF {
			for i := 1; i < sLen; i++ {
				if s[i] == sF {
					if isMatch(s[i+1:], p[2:]) {
						return true
					}
				} else {
					break
				}
			}
		}
		return false
	} else if  pF == '.' || sF == pF {
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}
