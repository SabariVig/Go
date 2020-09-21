package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		x = -x
		negative = true
	}
	var sum int
	for x > 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	if negative {
		sum = -sum

	}
	if sum > math.MaxInt32 {
		return 0

	}
	fmt.Println(sum)
	return sum
}

func main() {
	reverse(123)
}
