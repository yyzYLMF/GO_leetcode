package main

import "fmt"
import "sort"
import "math"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	max := math.MinInt32
	var i int
	for i = len(nums) - 1; i >= 0; i-- {
		if max > nums[i] {
			break
		} else {
			max = nums[i]
		}
	}
	if i < 0 {
		sort.Ints(nums)
	} else {
		min, min_pos := math.MaxInt32, -1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] && nums[j] < min {
				min = nums[j]
				min_pos = j
			}
		}
		nums[i], nums[min_pos] = nums[min_pos], nums[i]
		sort.Ints(nums[i+1:])
	}

	return
}

func main() {
	input := []int{16, 27, 25, 23, 25, 16, 12, 9, 1, 2, 7, 20, 19, 23, 16, 0, 6, 22, 16, 11, 8, 27, 9, 2, 20, 2, 13, 7, 25, 29, 12, 12, 18, 29, 27, 13, 16, 1, 22, 9, 3, 21, 29, 14, 7, 8, 14, 5, 0, 23, 16, 1, 20}
	input = []int{1, 3, 2}
	nextPermutation(input)
	fmt.Println(input)
	return
}
