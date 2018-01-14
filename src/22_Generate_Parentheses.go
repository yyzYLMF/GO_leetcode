package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}
	temp := ""
	ret := mydfs(temp, n, 0)
	return ret
}

func mydfs(temp string, left int, right int) []string {
	ret := []string{}
	if left == 0 {
		for i := 0; i < right; i++ {
			temp = temp + ")"
		}
		ret = append(ret, temp)
		return ret
	}
	addLeft := mydfs(temp+"(", left-1, right+1)
	ret = append(ret, addLeft...)
	if right > 0 {
		noLeft := mydfs(temp+")", left, right-1)
		ret = append(ret, noLeft...)
	}
	return ret
}

func main() {
	fmt.Println(generateParenthesis(3))
	return
}
