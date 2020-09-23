package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	refrence := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	for _, str := range s {
		if str == '{' || str == '(' || str == '[' {
			stack = append(stack, str)
		} else {
			if len(stack) == 0 || refrence[stack[len(stack)-1]] != str {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
