package p0166_Fraction_to_Recurring_Decimal

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	t.Log(fractionToDecimal(100, 3))
	t.Log(fractionToDecimal(100, 2))
	t.Log(fractionToDecimal(2, 3))
	t.Log(fractionToDecimal(2432, 5))
	t.Log(fractionToDecimal(-50, 8))
}
