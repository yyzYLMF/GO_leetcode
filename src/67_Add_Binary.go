package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	if len(a) < len(b) {
		return addBinary(b, a)
	}
	ret := ""
	tmp := 0
	i := len(a) - 1
	j := len(b) - 1
	for {
		if j < 0 {
			break
		}
		va := int(a[i]) - int('0')
		vb := int(b[j]) - int('0')
		ret = fmt.Sprintf("%+v", (va+vb+tmp)%2) + ret
		tmp = (va + vb + tmp) / 2
		i--
		j--
	}
	for {
		if i < 0 {
			break
		}
		va := int(a[i]) - int('0')
		ret = fmt.Sprintf("%+v", (va+tmp)%2) + ret
		tmp = (va + tmp) / 2
		i--
	}
	if tmp > 0 {
		ret = "1" + ret
	}
	return ret
}

func main() {
	fmt.Println(addBinary("11", "1"))
	return
}
