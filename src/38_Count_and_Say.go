package main

import "fmt"

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	tmp := "1"
	for i := 2; i <= n; i++ {
		last := tmp[0]
		num := 1
		ret := ""
		fmt.Println(tmp)
		for j := 1; j < len(tmp); j++ {
			if tmp[j] == last {
				num++
			} else {
				ret = fmt.Sprintf("%s%d%c", ret, num, last)
				last = tmp[j]
				num = 1
			}
		}
		ret = fmt.Sprintf("%s%d%c", ret, num, last)
		tmp = ret
	}
	return tmp
}

func main() {
	fmt.Println(countAndSay(4))
	return
}
