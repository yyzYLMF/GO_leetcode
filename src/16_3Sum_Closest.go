package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	div := math.MaxInt32
	ret := 0
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		t := target - nums[i]
		p, q := i+1, len(nums)-1
		for {
			if p >= q {
				break
			}
			if nums[p]+nums[q] == t {
				return target
			} else {
				if int(math.Abs(float64(nums[p]+nums[q]-t))) < div {
					div = int(math.Abs(float64(nums[p] + nums[q] - t)))
					ret = nums[i] + nums[p] + nums[q]
				}
				if nums[p]+nums[q] > t {
					q--
				} else {
					p++
				}
			}
		}
	}
	return ret
}

func main() {
	input := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(input, 1))
	return
}
