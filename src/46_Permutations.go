package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sign := make([]int, len(nums))
	ret := [][]int{}
	tmp := []int{}
	dfs(nums, tmp, sign, &ret, len(nums))
	return ret
}

func dfs(nums []int, tmp []int, sign []int, ret *[][]int, leftNum int) {
	if leftNum == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
	}
	for i, _ := range nums {
		if sign[i] == 1 {
			continue
		}
		sign[i] = 1
		tmp = append(tmp, nums[i])
		dfs(nums, tmp, sign, ret, leftNum-1)
		tmp = tmp[0 : len(tmp)-1]
		sign[i] = 0
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
	return
}
