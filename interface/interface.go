package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type rectangle struct {
	length float32
	width  float32
}

type square struct {
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func (c circle) area() float32 {
	return float32(math.Pi * math.Pow(float64(c.radius), 2))
}

type shape interface {
	area() float32
}

func main() {
	defaultStruct()

}

func defaultStruct() {
	s := square{65}
	r := rectangle{9, 6}
	c := circle{5}
	fmt.Println(s.area(), r.area(), c.area())
}
func usingInterface(){
	s := square{65}
	r := rectangle{9, 6}
	c := circle{5}
	shapes :=[]shape{s,r,c}
	for _,v := range shapes{
		fmt.Println(v.area())
	}

}