package p0168_Excel_Sheet_Column_Title

import "testing"

func Test_convertToTitle(t *testing.T) {
	t.Log(convertToTitle(701))
	t.Log(convertToTitle(703))
	t.Log(convertToTitle(28))
	t.Log(convertToTitle(26))
	t.Log(convertToTitle(52))
}
