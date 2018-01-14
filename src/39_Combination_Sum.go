package main

import "fmt"
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	tmp := []int{}
	mydfs(candidates, 0, &ret, tmp, target)
	return ret
}

func mydfs(nums []int, pos int, ret *[][]int, tmp []int, left int) {
	if left == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
		return
	}
	if pos >= len(nums) || left < 0 {
		return
	}
	tmp1 := append(tmp, nums[pos])
	mydfs(nums, pos, ret, tmp1, left-nums[pos])
	tmp2 := tmp
	mydfs(nums, pos+1, ret, tmp2, left)
	return
}

func main() {
	input := []int{2, 1}
	fmt.Println(combinationSum(input, 4))
	return
}
