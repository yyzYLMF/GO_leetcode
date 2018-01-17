package main

import "fmt"
import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	sign := make([]int, len(nums))
	tmp := []int{}
	ret := [][]int{}
	dfs(nums, sign, tmp, &ret, len(nums))
	return ret
}

func dfs(nums []int, sign []int, tmp []int, ret *[][]int, left int) {
	if left == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
		return
	}
	last := -1
	for i := 0; i < len(nums); i++ {
		if sign[i] == 1 {
			continue
		}
		if last != -1 && last != i && nums[last] == nums[i] {
			continue
		}
		sign[i] = 1
		tmp = append(tmp, nums[i])
		dfs(nums, sign, tmp, ret, left-1)
		tmp = tmp[0 : len(tmp)-1]
		sign[i] = 0
		last = i
	}
}

func main() {
	input := []int{1, 1, 2}
	fmt.Println(permuteUnique(input))
	return
}
