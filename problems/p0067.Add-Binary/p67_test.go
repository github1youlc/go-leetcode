package p67

import "testing"

func Test_addBinary(t *testing.T) {
	t.Log(addBinary("11", "1"))
	t.Log(addBinary("1010", "1011"))
	t.Log(addBinary("1010", ""))
}