package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	max := 0
	pos := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			if pos > 0 && stack[pos-1] == '(' {
				pos--
			} else {
				return false
			}
		case ']':
			if pos > 0 && stack[pos-1] == '[' {
				pos--
			} else {
				return false
			}
		case '}':
			if pos > 0 && stack[pos-1] == '{' {
				pos--
			} else {
				return false
			}
		default:
			if max > pos {
				stack[pos] = s[i]
				pos++
			} else {
				stack = append(stack, s[i])
				pos++
				max = pos
			}
		}
	}
	if pos == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	str := "(){}[()]"
	fmt.Println(isValid(str))
	return
}
