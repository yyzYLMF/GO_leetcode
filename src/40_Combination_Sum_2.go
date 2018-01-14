package main

import "fmt"
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	tmp := []int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	mydfs(&ret, tmp, candidates, 0, target)
	return ret
}

func mydfs(ret *[][]int, tmp []int, nums []int, pos int, left int) {
	if left == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
		return
	}

	if left < 0 {
		return
	}
	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		addTmp := append(tmp, nums[i])
		mydfs(ret, addTmp, nums, i+1, left-nums[i])
	}
	return
}

func main() {
	input := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(input, 8))
	return
}
