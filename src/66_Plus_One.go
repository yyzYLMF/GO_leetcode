package main

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	tmp := 1
	for i := len(digits) - 1; i >= 0; i-- {
		left := (digits[i] + tmp) % 10
		tmp = (digits[i] + tmp) / 10
		digits[i] = left
		if tmp == 0 {
			break
		}
	}
	if tmp == 1 {
		ret := []int{1}
		ret = append(ret, digits...)
		return ret
	} else {
		return digits
	}
}

func main() {
	return
}
