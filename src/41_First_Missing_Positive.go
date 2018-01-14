package main

import "fmt"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for pos, num := range nums {
		if num > len(nums) || num <= 0 {
			continue
		}
		t_pos := pos
		t_num := num
		for {
			if t_num == t_pos+1 || t_num > len(nums) || t_num <= 0 {
				break
			}
			nums[t_num-1], t_num, t_pos = t_num, nums[t_num-1], t_num-1
		}
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	input := []int{2, 1}
	fmt.Println(firstMissingPositive(input))
	return
}
