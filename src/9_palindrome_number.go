package main

import "fmt"

func isPalindrome(x int) bool {
	var tmp int64
	if x < 0 {
		return false
	} else {
		tmp = int64(x)
	}
	vec := []int64{}
	for {
		if tmp == 0 {
			break
		}
		vec = append(vec, tmp%10)
		tmp = tmp / 10
	}
	var i, j int
	for i, j = 0, len(vec)-1; i < j; i, j = i+1, j-1 {
		if vec[i] != vec[j] {
			break
		}
	}
	if i < j {
		return false
	} else {
		return true
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(isPalindrome(n))
	return
}
