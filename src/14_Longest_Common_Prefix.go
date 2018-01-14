package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	tmp := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			tmp = strs[i]
		}
	}

	for i := 0; i < len(strs); i++ {
		tmp = myfunc(tmp, strs[i])
	}
	return tmp
}

func myfunc(str1, str2 string) string {
	ret := 0
	for i, j := 0, 0; i < len(str1) && j < len(str2); i, j = i+1, j+1 {
		if str1[i] != str2[j] {
			break
		}
		ret++
	}
	return str1[:ret]
}

func main() {
	strs := []string{"hello", "he", "hell"}
	fmt.Println(longestCommonPrefix(strs))
	return
}
