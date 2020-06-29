package p49

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type strIdx struct {
	s string
	i int
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	cp := make([]*strIdx, len(strs))
	for i, s := range strs {
		cp[i] = &strIdx{
			s: SortString(s),
			i: i,
		}
	}

	ret := make([][]string, 0)
	m := make(map[string][]int)
	for _, si := range cp {
		m[si.s] = append(m[si.s], si.i)
	}

	for _, idxs := range m {
		arr := make([]string, 0, len(idxs))
		for _, idx := range idxs {
			arr = append(arr, strs[idx])
		}
		ret = append(ret, arr)
	}

	return ret
}
