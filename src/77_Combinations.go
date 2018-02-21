package main

import "fmt"

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}
	ret := [][]int{}
	tmp := []int{}
	mydfs(&ret, tmp, 1, n, k)
	return ret
}

func mydfs(ret *[][]int, tmp []int, pos, n, left int) {
	if left == 0 {
		item := []int{}
		item = append(item, tmp...)
		*ret = append(*ret, item)
		return
	}
	if pos > n {
		return
	}
	tmp2 := append(tmp, pos)
	mydfs(ret, tmp2, pos+1, n, left-1)
	mydfs(ret, tmp, pos+1, n, left)
	return
}

func main() {
	fmt.Println(combine(4, 2))
	return
}
