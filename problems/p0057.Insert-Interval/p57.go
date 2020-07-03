package p57

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int , 0)
	n := len(intervals)
	pos := 0

	start := newInterval[0]
	end := newInterval[1]

	for ; pos < n; pos ++ {
		if start < intervals[pos][0] {
			break
		}
	}


	left := pos -1
	for ; left >=0; left -- {
		v := intervals[left]
		if v[1] < start {
			break
		} else {
			start = v[0]
			end = max(end, v[1])
		}
	}


	right := pos
	for ; right < n; right ++ {
		v := intervals[right]
		if v[0] > end {
			break
		} else {
			end = max(end, v[1])
		}
	}

	if left >= 0 {
		ret = append(ret, intervals[0:left+1]...)
	}

	ret = append(ret, []int{start, end})

	if right < n {
		ret = append(ret, intervals[right:n]...)
	}

	return ret

}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}