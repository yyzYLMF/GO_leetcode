package main

import (
	"bufio"
	"fmt"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	store := map[byte]int{}
	ret := 0
	start := 0
	tmpLen := 0
	tmpStr := ""
	for i := 0; i < len(s); i = i + 1 {
		val, ok := store[s[i]]
		if !ok || val < start {
			store[s[i]] = i
			tmpStr = tmpStr + string(s[i])
			tmpLen += 1
			if tmpLen > ret {
				ret = tmpLen
			}
		} else {
			tmpStr = s[val+1 : i+1]
			tmpLen = len(tmpStr)
			store[s[i]] = i
			start = val + 1
		}
	}
	return ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	str := string(input[0 : len(input)-1])
	fmt.Println(str) // string(input[0:len(input)-1]) remove '\n'.
	fmt.Println(lengthOfLongestSubstring(str))
	return
}
