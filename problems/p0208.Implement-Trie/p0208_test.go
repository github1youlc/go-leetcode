package solve

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Insert("awordz")
	obj.Insert("aword")
	t.Log(obj.Search("awordz"))
	t.Log(obj.Search("aword"))
	t.Log(obj.Search("awor"))
	t.Log(obj.StartsWith("awo"))
	t.Log(obj.StartsWith("b"))
}
