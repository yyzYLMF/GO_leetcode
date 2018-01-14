package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	mm := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mm[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			if _, ok := mm[target-nums[j]]; ok {
				if mm[target-nums[j]] > j {
					item := []int{nums[i], nums[j], target - nums[j]}
					ret = append(ret, item)
				}
			}
		}
	}
	return ret
}

func main() {
	input := []int{-3, 14, -10, -1, 12, 13, -3, 2, -6, 4, 13, 7, -8, 4, 0, -13, 11, -4, 7, 0, 4, -3, 12, 11, 5, -14, -8, 8, 3, -1, -8, -15, -2, -11, -9, -12, 9, 3, 5, -11, -8, 3, 3, -9, -15, -12, -15, 3, -9, 0, -12, 3, 12, -14, -1, -6, -13, -2, -13, -3, 12, -14, -3, -13, -9, 3, -10, -15, 13, 2, 11, 13, -9, -1, 11, 13, -6, 4, 1, 1, -11, 5, -11, 8, -2, -5, -12, -8, 8, -10, 4, -3, -8, -14, -1, -10, -4, -3, 12, -14, 14, 9, 6, 12, -15, 3, 10, -13, -8, -11, 3, -4, 1, -1}
	fmt.Println(threeSum(input))
	return
}
