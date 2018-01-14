package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1
	ret := -1
	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			high = mid - 1
			ret = mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ret
}

func findLast(nums []int, target int) int {
	low, high := 0, len(nums)-1
	ret := -1
	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			low = mid + 1
			ret = mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ret
}

func main() {
	input := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(input, 8))
	return
}
