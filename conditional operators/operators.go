package main

import "fmt"

func main() {
	x, y := 5, 6
	val := x <= y
	fmt.Println(val)
	alpha()
}

func alpha() {
	x, y := "B", "A"
	val := x <= y || 4 < 5 
	fmt.Println(val)
}
