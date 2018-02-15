package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	tmp := make([]int, n+1)
	tmp[0], tmp[1] = 1, 1
	for i := 2; i <= n; i++ {
		tmp[i] = tmp[i-1] + tmp[i-2]
	}
	return tmp[n]
}

func main() {
	fmt.Println(climbStairs(4))
	return
}
