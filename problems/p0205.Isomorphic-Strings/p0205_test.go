package solve

import "testing"

func Test_isIsomorphic(t *testing.T) {
	t.Log(isIsomorphic("abc", "abd"))
	t.Log(isIsomorphic("app", "egg"))
	t.Log(isIsomorphic("foo", "bar"))
	t.Log(isIsomorphic("bar", "foo"))
}
