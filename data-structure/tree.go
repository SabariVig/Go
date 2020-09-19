package main

import "fmt"

type tree struct {
	lnode *tree
	rnode *tree
	key   int
}

func (t *tree) insert(k int) {
	if t.key > k {
		if t.lnode == nil {
			t.lnode = &tree{key: k}
		} else {
			t.lnode.insert(k)
		}
	}
	if t.key < k {
		if t.rnode == nil {
			t.rnode = &tree{key: k}
		} else {
			t.rnode.insert(k)
		}

	}
}

func (t *tree) search(k int) {
	fmt.Println(t.key)
	if t == nil {
		fmt.Print("Not Found")
	}
	if t.key < k {
		t.rnode.search(k)
	} else if t.key > k {
		t.lnode.search(k)
	}
}

func main() {
	tre := &tree{key: 5}
	tre.insert(28)
	tre.insert(20)
	tre.insert(10)
	tre.insert(3)
	tre.search(10)
	// fmt.Println(tre)
}
