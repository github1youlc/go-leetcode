package solve

import "testing"

func TestWordDictionary_Search(t *testing.T) {
	wd := Constructor()
	wd.AddWord("abc")
	wd.AddWord("dad")
	wd.AddWord("bad")
	t.Log(wd.Search(".bc"))
	t.Log(wd.Search("bag"))
	t.Log(wd.Search("ab"))
	t.Log(wd.Search(".b"))
	t.Log(wd.Search("abc"))
}