package p51

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	{
		ret := solveNQueens(4)
		for _, r := range ret {
			fmt.Println(r)
		}
	}
}