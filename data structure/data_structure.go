package main

import "fmt"

func main() {
	arr := array()
	for i, v := range arr {
		fmt.Println("Index:", i, "Value:", v)
	}
	maps := maps()

	for k, v := range maps {
		fmt.Println("Key:", k, "Value:", v)
	}

}
func array() ([]int,error) {
	x := []int{2, 5}
	x = append(x, 4)
	// fmt.Println("Array:", x)
	return x, nil 

}

func maps() map[string]int8 {
	vertices := make(map[string]int8)
	vertices["Triangle"] = 3
	vertices["Square"] = 4
	vertices["Pentagon"] = 5

	// fmt.Scan( &vertices["City"] )
	// fmt.Println("Map:", vertices)
	return vertices

}
