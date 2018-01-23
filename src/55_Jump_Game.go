package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	pos := 0
	for {
		maxStep := nums[pos]
		if pos >= len(nums)-1 {
			return true
		}
		if maxStep == 0 {
			return false
		}
		maxNext := 0
		nextPos := 0
		for i := 1; i <= maxStep; i++ {
			if pos+i >= len(nums)-1 {
				return true
			}
			if pos+i+nums[pos+i] > maxNext {
				maxNext = pos + i + nums[pos+i]
				nextPos = pos + i
			}
		}
		pos = nextPos
		fmt.Println("next ", nextPos, maxNext)
	}
	return false
}

func main() {
	input := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println(canJump(input))

	return
}
