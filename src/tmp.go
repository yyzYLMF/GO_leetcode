package main

import "fmt"
import "sort"

func main() {
	input := []int{1, 3, 5, 2, 4, 6}
	input2 := append(input, 3)
	fmt.Println("1", &input2[0], input2)
	myfunc(input2)
	fmt.Println("3", &input2[0], len(input2), input2)
	return
}

func myfunc(input []int) {
	input = append(input, 6)
	sort.Ints(input)
	fmt.Println("2", &input[0], len(input), input)
	var c = input
	fmt.Println(&c[0])
	return
}
