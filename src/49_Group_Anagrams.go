package main

import "fmt"
import "sort"

type ByteSlice []byte

func (this ByteSlice) Len() int {
	return len(this)
}
func (this ByteSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this ByteSlice) Less(i, j int) bool {
	return this[i] < this[j]
}

func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}
	for _, str := range strs {
		byteStr := []byte(str)
		sort.Sort(ByteSlice(byteStr))
		key := string(byteStr)
		if _, ok := strMap[key]; !ok {
			strMap[key] = []string{}
		}
		strMap[key] = append(strMap[key], str)
	}
	ret := [][]string{}
	for _, value := range strMap {
		ret = append(ret, value)
	}
	return ret
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
	return
}
