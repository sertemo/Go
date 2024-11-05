package main

import "fmt"

type Tree struct {
	Left, Right *Tree
	Value       int
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.Left)
		values = append(values, t.Value)
		values = appendValues(values, t.Right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{value, nil, nil}
		t = new(Tree)
		t.Value = value
		return t
	}
	if value < t.Value {
		t.Left = add(t.Left, value)
	} else {
		t.Right = add(t.Right, value)
	}
	return t
}

func main() {
	values := []int{12, 2, 3, 73, 25, 66, 17, 8, 9, 10}
	Sort(values)
	fmt.Println(values)
}
