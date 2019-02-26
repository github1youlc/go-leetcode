/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-2-25 下午1:13
 */

package p0347_Top_K_Frequent_Elements

import "container/heap"

type countItem struct {
	count int
	num   int
}

type countHeap []*countItem

func (h countHeap) Len() int           { return len(h) }
func (h countHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h countHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *countHeap) Push(x interface{}) {
	*h = append(*h, x.(*countItem))
}

func (h *countHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	countMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			countMap[num]++
		} else {
			countMap[num] = 1
		}
	}

	h := &countHeap{}
	heap.Init(h)
	for num, count := range countMap {
		if h.Len() < k {
			heap.Push(h, &countItem{
				count: count,
				num:   num,
			})
		} else {
			if (*h)[0].count < count {
				heap.Pop(h)
				heap.Push(h, &countItem{
					count: count,
					num:   num,
				})
			}
		}
	}

	result := make([]int, h.Len())
	cur := h.Len() - 1
	for h.Len() > 0 {
		item := heap.Pop(h)
		result[cur] = item.(*countItem).num
		cur--
	}

	return result
}
