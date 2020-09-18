package solve

import "container/heap"

type intHeap []int

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *intHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return 0
	}

	h := &intHeap{

	}
	heap.Init(h)
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}

	return heap.Pop(h).(int)
}