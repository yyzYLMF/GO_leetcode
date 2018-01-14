package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mvalue := map[int]int{}
	for index, value := range nums {
		left := target - value
		if _, ok := mvalue[left]; ok {
			return []int{mvalue[left], index}
		} else {
			mvalue[value] = index
		}
	}
	return nil
}

func main() {
	input := []int{1, 2, 3, 6}
	ret := twoSum(input, 5)
	fmt.Println(ret)
}
