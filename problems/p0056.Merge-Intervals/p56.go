package p56

import "sort"

type intervalSort [][]int

func (s intervalSort) Len() int {
	return len(s)
}

func (s intervalSort) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s intervalSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Sort(intervalSort(intervals))

	last := intervals[0]
	ret := make([][]int, 0)

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]

		if last[1] < cur[0] {
			ret = append(ret, last)
			last = cur
		} else {
			last[1] = max(last[1], cur[1])
		}
	}

	ret = append(ret, last)

	return ret
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}