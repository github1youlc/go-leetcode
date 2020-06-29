package p17

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	a := [][]byte{{1, 2,}}
	b := []byte{5, 6,7}


	s := make([][]byte, 0, len(a) * len(b))
	for i:=0; i<3; i++ {
		s = multiply(a, b)
		a = s
	}

	for _, si := range s {
		fmt.Println(si)
	}
}