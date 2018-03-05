package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	pos := 0
	lastNum := 1
	lastValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastValue {
			if lastNum < 2 {
				nums[pos+1] = nums[i]
				pos++
				lastNum++
			}
		} else {
			nums[pos+1] = nums[i]
			pos++
			lastNum = 1
			lastValue = nums[i]
		}
	}
	return pos + 1
}

func main() {
	input := []int{1, 1, 1, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(input))
	return
}
