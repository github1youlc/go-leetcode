package p0152_Maximum_Product_SubArray

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	p := product(nums)
	if p > 0 {
		return p
	}

	prev := -1
	max := p
	hasZero := false
	for i := 0; i < n; i++ {
		for i < n && nums[i] != 0 {
			i++
		}

		if i < n {
			hasZero = true
		}

		if prev < i {
			fmt.Println(prev+1, ",", i)
			p = maxProductNoZero(nums[prev+1:i])
			if p > max {
				max = p
			}
		}

		prev = i

		if i == n {
			break
		}
	}

	if hasZero && max < 0 {
		max = 0
	}

	return max
}


func maxProductNoZero(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	p := product(nums)

	if p >= 0 {
		return p
	}

	var left, right int
	for i, num := range nums {
		if num < 0 {
			left = i
			break
		}
	}

	for i := n -1 ;i >=0; i-- {
		if nums[i] < 0 {
			right = i
			break
		}
	}

	return maxInt(product(nums[0:right]), product(nums[left+1: n]))
}

func product(nums []int) int {
	p := 1
	for _, n := range nums {
		p *= n
	}

	return p
}

func maxInt(a int,  b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}