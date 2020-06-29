/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 18-12-31 下午4:20
 */

package p30

import (
	"testing"
)

func Test_findSubstring(t *testing.T) {
	t.Log(findSubstring("barfoothefoobarman", []string{"foo","bar"}))
	t.Log(findSubstring("wordgoodgoodgoodbestword",
	[]string{"word","good","best","word"}))
}
