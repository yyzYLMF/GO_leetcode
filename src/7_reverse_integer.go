package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var tmp, ret int64
	flag := false
	if x == 0 {
		return x
	}
	if x < 0 {
		flag = true
		tmp = 0 - int64(x)
	} else {
		tmp = int64(x)
	}
	ret = 0
	for {
		if tmp == 0 {
			break
		}
		ret = ret*10 + tmp%10
		tmp = tmp / 10
	}
	if flag {
		ret = 0 - ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	} else {
		return int(ret)
	}
}

func main() {
	var input int
	fmt.Scanf("%d", &input)
	ret := reverse(input)
	fmt.Println(ret)
	return
}
