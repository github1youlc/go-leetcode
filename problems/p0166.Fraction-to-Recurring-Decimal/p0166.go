package p0166_Fraction_to_Recurring_Decimal

import "strconv"

type history struct {
	result    []int
	traversed map[int]int
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator > 0 && denominator > 0 {
		return fractionToDecimalInner(numerator, denominator)
	} else if numerator > 0 && denominator < 0 {
		return "-" + fractionToDecimalInner(numerator, -denominator)
	} else if numerator < 0 && denominator > 0 {
		return "-" + fractionToDecimalInner(-numerator, denominator)
	} else if numerator < 0 && denominator < 0 {
		return fractionToDecimalInner(-numerator, -denominator)
	} else {
		return "0"
	}
}

func fractionToDecimalInner(numerator int, denominator int) string {
	nn := numerator / denominator
	nr := numerator % denominator

	if nr == 0 {
		return strconv.Itoa(nn)
	} else {
		h := &history{
			traversed: make(map[int]int),
		}
		repeatBegin := -1
		decimal(nr, denominator, h, &repeatBegin)

		return strconv.Itoa(nn) + "." + history2StringWithRepeat(h, repeatBegin)
	}
}

func history2StringWithRepeat(h *history, repeatBegin int) string {
	if repeatBegin == -1 {
		return history2String(h.result)
	} else {
		return history2String(h.result[0:repeatBegin]) + "(" + history2String(h.result[repeatBegin:]) + ")"
	}
}

func history2String(res []int) string {
	var result string
	for i := 0; i < len(res); i++ {
		result += strconv.Itoa(res[i])
	}
	return result
}

func decimal(n int, d int, h *history, repeatBegin *int) {
	id := generateId(n, d)
	hisPos := inHistory(id, h)

	if hisPos != -1 {
		*repeatBegin = hisPos
		return
	}

	nn := n * 10 / d
	nr := n * 10 % d

	h.traversed[id] = len(h.traversed)
	h.result = append(h.result, nn)

	if nr != 0 {
		decimal(nr, d, h, repeatBegin)
	}
}

func inHistory(id int, h *history) int {
	if idx, ok := h.traversed[id]; ok {
		return idx
	} else {
		return -1
	}
}

func generateId(n int, d int) int {
	return n*10 + d
}
