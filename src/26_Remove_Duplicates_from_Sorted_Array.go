package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return len(nums)
	}
	p, q := 0, 1
	last := nums[p]
	for ; q < len(nums); q++ {
		if last != nums[q] {
			p++
			nums[p] = nums[q]
			last = nums[q]
		}
	}
	return p + 1
}

func main() {
	input := []int{1, 1, 2}
	fmt.Println(removeDuplicates(input))
	return
}
