package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ret := [][]int{}
	for i := 0; i <= len(nums); i++ {
		item := [][]int{}
		tmp := []int{}
		mydfs(&item, nums, 0, tmp, i)
		ret = append(ret, item...)
	}
	return ret
}

func mydfs(ret *[][]int, nums []int, pos int, tmp []int, left int) {
	if left == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
		return
	}
	if pos >= len(nums) {
		return
	}
	tmp2 := append(tmp, nums[pos])
	mydfs(ret, nums, pos+1, tmp2, left-1)
	mydfs(ret, nums, pos+1, tmp, left)
	return
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	return
}
