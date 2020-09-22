package main

import "fmt"

func isPalindrome(x int) bool {
	org := x
	reverse := 0
	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	if reverse != org {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
