package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for {
		if low > high {
			return -1
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[mid] < nums[low] {
				high = mid - 1
			} else {
				if nums[low] <= target {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
		} else {
			if nums[mid] > nums[high] {
				low = mid + 1
			} else {
				if nums[high] >= target {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		}
	}
	return -1
}

func main() {
	input := []int{4, 5, 6, 7, 8, 1, 2, 3}
	fmt.Println(search(input, 8))
	return
}
