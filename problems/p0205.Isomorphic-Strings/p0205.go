package solve

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[uint8]uint8)
	rm := make(map[uint8]uint8)

	var sc uint8
	var tc uint8
	for i := range s {
		sc = s[i]
		tc = t[i]

		if ms, ok := m[sc]; ok {
			if ms != tc {
				return false
			}

			if rmt, ok := rm[tc]; !ok || rmt != sc {
				return false
			}
		} else {
			if _, ok := rm[tc]; ok {
				return false
			}

			m[sc] = tc
			rm[tc] = sc
		}
	}

	return true
}