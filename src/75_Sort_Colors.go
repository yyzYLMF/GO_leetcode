package main

import "fmt"

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	p, q := -1, len(nums)
	for i := 0; i < q; {
		if nums[i] == 0 {
			nums[p+1], nums[i] = nums[i], nums[p+1]
			p++
			i++
		} else if nums[i] == 2 {
			nums[q-1], nums[i] = nums[i], nums[q-1]
			q--
		} else {
			i++
		}
	}
	return
}

func main() {
	input := []int{1, 0, 2, 2, 2}
	sortColors(input)
	fmt.Println(input)
	return
}
