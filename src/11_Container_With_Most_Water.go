package main

import "fmt"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	p, q := 0, len(height)-1
	max := 0
	for {
		if p >= q {
			break
		}
		var tmp int
		if height[p] >= height[q] {
			tmp = (q - p) * height[q]
			q--
		} else {
			tmp = (q - p) * height[p]
			p++
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	height := []int{1, 1}
	fmt.Println(maxArea(height))
	return
}
