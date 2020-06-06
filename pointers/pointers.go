package main

import "fmt"

func changeRef(a *int) {
	*a += 2
}

func changeVal(a int) {
	a += 2
}

func main() {
	x := 55
	fmt.Println(x)
	changeRef(&x)
	fmt.Println(x) //Prints The Modified Value Of X
	y := 22
	fmt.Println(y)
	changeVal(y)
	fmt.Println(y) // Prints Same Value As Y

}
