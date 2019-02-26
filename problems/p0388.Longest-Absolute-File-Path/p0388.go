/*
Suppose we abstract our file system by a string in the following manner:


The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:

dir
subdir1
subdir2
file.ext
The directory dir contains an empty sub-directory subdir1 and a sub-directory subdir2 containing a file file.ext.

The string "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" represents:

dir
subdir1
file1.ext
subsubdir1
subdir2
subsubdir2
file2.ext
The directory dir contains two sub-directories subdir1 and subdir2. subdir1 contains a file file1.ext and an empty second-level sub-directory subsubdir1. subdir2 contains a second-level sub-directory subsubdir2 containing a file file2.ext.

We are interested in finding the longest (number of characters) absolute path to a file within our file system. For example, in the second example above, the longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and its length is 32 (not including the double quotes).

Given a string representing the file system in the above format, return the length of the longest absolute path to file in the abstracted file system. If there is no file in the system, return 0.

Note:
The name of a file contains at least a . and an extension.
The name of a directory or sub-directory will not contain a ..
Time complexity required: O(n) where n is the size of the input string.

Notice that a/aa/aaa/file1.txt is not the longest file path, if there is another path aaaaaaaaaaaaaaaaaaaaa/sth.png.

 */

package p0338

import (
	"bytes"
	"fmt"
	"strings"
)

type dir struct {
	parent  *dir
	pathLen int
	depth   int
	path 	string
}

func (d *dir) seekToDepth(depth int) *dir {
	delta := d.depth - depth
	ret := d
	for i := 0; i < delta; i++ {
		ret = ret.parent
	}

	return ret
}

func lengthLongestPath(input string) int {
	root := &dir{
		parent:  nil,
		depth:   -1,
		pathLen: 0,
		path: "",
	}

	cur := root

	maxLen := 0
	var length int
	var item string
	for i := 0; i < len(input); {
		c := input[i]
		depth := 0
		if c == '\t' {
			depth = 1
			for j := i + 1; j < len(input); j++ {
				if input[j] == '\t' {
					depth++
				} else {
					i = j
					break
				}
			}
		}

		parent := cur.seekToDepth(depth - 1)
		item, i = nextItem(input, i)

		if isFile(item) {
			fmt.Println(parent.path + item)
			length = parent.pathLen + len(item)
			if length > maxLen {
				maxLen = length
			}
			cur = parent
		} else {
			cur = &dir{
				parent:  parent,
				pathLen: parent.pathLen + len(item) + 1,
				depth:   parent.depth + 1,
				path: parent.path + item + "/",
			}
		}
	}
	return maxLen
}

func isFile(s string) bool {
	return strings.Contains(s, ".")
}

func nextItem(input string, i int) (string, int) {
	var buf bytes.Buffer
	var j int
	for j = i; j < len(input); j++ {
		if input[j] != '\n' {
			buf.WriteByte(input[j])
		} else {
			break
		}
	}

	return buf.String(), j + 1
}
