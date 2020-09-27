package main

import "fmt"

type Node struct {
	key  int
	next *Node
}
type List struct {
	head   *Node
	length int
}

func (l *List) Prepend(val int) {
	tmp := &Node{key: val}
	tmp.next = l.head
	l.head = tmp
	l.length += 1
}

func (l *List) Append(val int) {
	tmp := &Node{key: val}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = tmp
	l.length += 1
}
func (l *List) Delete(val int) {
	toDel := l.head
	for toDel.next.key != val {

	}

}

func (l List) Print() {
	for l.length != 0 {
		fmt.Println(l.head.key)
		l.length -= 1
		l.head = l.head.next
	}
}

func main() {
	var ll List
	ll.Prepend(10)
	ll.Prepend(20)
	ll.Prepend(70)
	ll.Prepend(50)
	ll.Prepend(20)
	ll.Prepend(22)
	ll.Append(8)
	ll.Print()
}
