package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	p, q := 0, 0
	for ; q < len(nums); q++ {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
		}
	}
	return p
}

func main() {
	input := []int{3, 2, 2, 3}
	fmt.Println(removeElement(input, 3))
	fmt.Println(1000000 / 1e6)
	return
}
