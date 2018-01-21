package main

import "fmt"
import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, tmp := math.MinInt32, 0
	for i := 0; i < len(nums); i++ {
		tmp = tmp + nums[i]
		if tmp > maxSum {
			maxSum = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return maxSum
}

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input))
	return
}
