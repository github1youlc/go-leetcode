package p38

import "testing"

func Test_countAndSay(t *testing.T) {
	t.Log(countAndSay(2))
	for i := 1; i < 20; i++ {
		t.Log(countAndSay(i))
	}

}