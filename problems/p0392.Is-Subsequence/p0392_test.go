package solve

import "testing"

func Test_isSubsequence(t *testing.T) {
	t.Log(isSubsequence("abc", "ahbgdc"))
	t.Log(isSubsequence("axc", "ahbgdc"))
}
