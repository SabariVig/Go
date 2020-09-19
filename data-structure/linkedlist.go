package main

import "fmt"

type node struct {
	data int
	next *node
}
type link struct {
	head   *node
	length int
}

func (l *link) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length += 1
}
func (l link) print() {
	fmt.Println(l)
	for l.length != 0 {
		fmt.Print(l.head.data, " ")
		l.length -= 1
		l.head = l.head.next
	}
	fmt.Print("\n")
}
func (l *link) delet(value int) {
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
func main() {
	var l link
	l.prepend(&node{data: 3})
	l.prepend(&node{data: 5})
	l.prepend(&node{data: 32})
	l.prepend(&node{data: 55})
	l.prepend(&node{data: 22})
	l.print()
	l.delet(32)
	l.print()

}
