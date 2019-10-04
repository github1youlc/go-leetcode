package solve

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	vs1 := strings.Split(version1, ".")
	vs2 := strings.Split(version2, ".")

	len1 := len(vs1)
	len2 := len(vs2)
	maxLen := maxInt(len1, len2)

	i := 0
	var v1 string
	var v2 string
	for i < maxLen {
		if i >= len1 {
			v1 = ""
		} else {
			v1 = vs1[i]
		}

		if i >= len2 {
			v2 = ""
		} else {
			v2 = vs2[i]
		}

		cv := compare(v1, v2)

		if cv == -1 {
			return -1
		} else if cv == 1 {
			return 1
		}

		i++
	}

	return 0
}


func compare(ver1 string, ver2 string) int {
	v1 := preProcess(ver1)
	v2 := preProcess(ver2)

	if v1 < v2 {
		return -1
	} else if v1 == v2 {
		return 0
	} else {
		return 1
	}
}

func preProcess(v string) int {
	var i int
	for i = 0; i < len(v); i++ {
		if v[i] != '0' {
			break
		}
	}

	v = v[i:]
	if v == "" {
		return 0
	}

	vInt, _ :=  strconv.Atoi(v)
	return vInt
}


func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}