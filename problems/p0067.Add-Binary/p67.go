package p67

func addBinary(a string, b string) string {
	na := len(a)
	nb := len(b)

	nm := max(na, nb)

	ret := make([]byte, nm + 1)

	var ai, bi int
	var av, bv uint8
	var carry uint8 = 0
	var sum uint8

	for i := 0 ; i < nm; i++ {
		ai = na - 1 - i
		bi = nb - 1 - i

		if ai < 0 {
			av = 0
		} else {
			av = a[ai] - '0'
		}

		if bi < 0 {
			bv = 0
		} else {
			bv = b[bi] - '0'
		}

		sum = av + bv + carry
		carry = sum / 2
		ret[nm-i] = sum % 2 + '0'
	}

	if carry == 0 {
		return string(ret[1:])
	} else {
		ret[0] = '1'
		return string(ret)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}