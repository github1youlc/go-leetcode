package p39

import "sort"

type intSort []int

func (s intSort) Len() int {
	return len(s)
}

func (s intSort) Swap(i int, j int)   {
	s[i], s[j] = s[j], s[i]
}

func (s intSort) Less(i int, j int) bool {
	return s[i] < s[j]
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Sort(intSort(candidates))

	return doCombinationSum(candidates, 0, nil, target)
}

func doCombinationSum(candidates []int, start int, buf []int,  target int) [][]int {
	if target == 0 {
		dst := make([]int, len(buf))
		copy(dst, buf)
		return [][]int{
			dst,
		}
	}

	if target < 0 {
		return nil
	}

	ret := make([][]int, 0)
	for i := start; i < len(candidates); i++ {
		c := candidates[i]
		if target >= c {
			buf = append(buf, c)
			ret = append(ret, doCombinationSum(candidates, i, buf, target - c)...)
			buf = buf[0: len(buf) - 1]
		} else {
			break
		}
	}

	return ret
}