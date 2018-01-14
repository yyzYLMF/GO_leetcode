package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	input := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(input, 5))
	return
}
