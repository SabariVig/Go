package main

import "fmt"

func main() {

	a, b := 15, 95

	c := a + b
	d := a - b
	e := a * b
	f := a / b
	x, y := 5.6, 7
	z := x / float64(y)

	fmt.Println(a, b, c, d, e, f, z)

}
