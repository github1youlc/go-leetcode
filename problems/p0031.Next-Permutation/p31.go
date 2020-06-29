// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
//If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
//
//The replacement must be in-place and use only constant extra memory.
//
//Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
//
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

package p31

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)

	i := n -1
	for ; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}

	if i == 0 {
		reverse(nums)
	} else {
		swapPos := i - 1
		swapNum := nums[swapPos]

		min := nums[i]
		minPos := i
		for j := i+1 ; j < n; j++ {
			if nums[j] > swapNum && nums[j] <= min {
				min = nums[j]
				minPos = j
			}
		}
		swap(nums, swapPos, minPos)
		reverse(nums[i:])
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n / 2; i++ {
		swap(nums, i, n -1 -i)
	}
}
