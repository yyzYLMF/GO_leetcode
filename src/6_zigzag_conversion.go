package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}
	tmp := make([]string, numRows)
	k := 2*numRows - 2
	for i := 0; i < len(s); i = i + 1 {
		pos := i % (k)
		if pos < numRows {
			tmp[pos] = tmp[pos] + string(s[i])
		} else {
			index := numRows - 2 - pos%numRows
			tmp[index] = tmp[index] + string(s[i])
		}
	}
	var ret string
	for i := 0; i < len(tmp); i = i + 1 {
		ret = ret + tmp[i]
	}
	return ret
}

func main() {
	fmt.Println("succ")
	return
}
