package p34

func searchRange(nums []int, target int) []int {
	n := len(nums)

	s := 0
	e := n - 1
	var m int
	var mid int
	found := false

	for s <= e {
		m = (s + e) / 2
		mid = nums[m]

		if mid == target {
			found = true
			break
		} else if target < mid {
			e = m - 1
		} else {
			s = m + 1
		}
	}

	if !found {
		return []int{-1, -1}
	}

	// left
	ret := make([]int, 0, 2)

	var cur int
	for cur = m - 1; cur >= 0 && nums[cur] == target; cur-- {
	}
	ret = append(ret, cur+1)

	for cur = m + 1; cur < n && nums[cur] == target; cur++ {
	}
	ret = append(ret, cur-1)

	return ret
	//matchPos := m
	//s = 0
	//e = matchPos

	//for s <= e {
	//	if s == e {
	//		ret = append(ret, s)
	//		break
	//	}
	//	m = (s + e) / 2
	//	mid = nums[m]
	//
	//	if mid == target {
	//		e = m
	//	} else {
	//		s = m + 1
	//	}
	//}
	//
	//s = matchPos
	//e = n - 1
	//for s <= e {
	//	if s == e {
	//		ret = append(ret, s)
	//		break
	//	}
	//
	//	m = (s + e + 1) / 2
	//	mid = nums[m]
	//
	//	if mid == target {
	//		s = m
	//	} else {
	//		e = m - 1
	//	}
	//}

	//return ret
}
