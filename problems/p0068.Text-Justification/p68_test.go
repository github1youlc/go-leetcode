package p68

import (
	"testing"
)

func Test_fullJustify(t *testing.T) {
	{
		words := []string {
			"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do",
		}

		ret := fullJustify(words, 20)
		for _, s := range ret {
			t.Log("|", s, "|")
		}
	}
}