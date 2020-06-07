package main

import "fmt"

type student struct {
	name    string
	marks   []int
	average float32
}
type account struct {
	id          int
	transaction []int
	average     float32
	sum         float32
}

func (s *student) computeAverage() {
	sum := 0
	for _, v := range s.marks {
		sum += v
	}
	s.average = float32(sum) / float32(len(s.marks))
}

func (a *account) computeAverage() {
	sum := 0
	for _, v := range a.transaction {
		sum += v
	}
	a.average = float32(sum) / float32(len(a.transaction))
	a.sum = float32(sum)

}

type person interface {
	computeAverage()
}

func main() {

	s := student{name: "Alex", marks: []int{80, 89, 98, 99, 100, 96}}
	// fmt.Println(s)
	// s.computeAverage()
	// fmt.Println(s)
	a := account{id: 5, transaction: []int{5, 10, 50, 10}}

	p := []person{&s, &a}

	fmt.Println(s,a)
	for _, v := range p {
		v.computeAverage()
	}

	fmt.Println(s,a)

}
