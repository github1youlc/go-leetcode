package p0155_Min_Stack

type MinStack struct {
	data     []int
	n        int
	capacity int
}

func newMinStack() *MinStack {
	return &MinStack{
		capacity: 2,
		data:     make([]int, 2),
		n:        0,
	}
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		capacity: 2,
		data:     make([]int, 2),
		n:        0,
	}
}

func (this *MinStack) Push(x int) {
	if this.n >= this.capacity {
		this.data = append(this.data, make([]int, this.capacity)...)
		this.capacity *= 2
	}

	this.data[this.n] = x
	this.n++
}

func (this *MinStack) Pop() {
	this.n--
}

func (this *MinStack) Top() int {
	if this.n <= 0 {
		return 0
	}

	return this.data[this.n - 1]
}

func (this *MinStack) GetMin() int {
	if this.n == 0 {
		return 0
	}

	min := this.data[0]
	for i := 1; i < this.n; i++ {
		if this.data[i] < min {
			min = this.data[i]
		}
	}

	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
