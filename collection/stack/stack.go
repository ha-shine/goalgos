package main

import (
	"fmt"
	"math/rand"
)

type stackNode struct {
	Item interface{}
	Next *stackNode
}

type Stack struct {
	head *stackNode
	Size int
}

func NewStack() *Stack {
	return &Stack{head: nil, Size: 0}
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size == 0
}

func (stack *Stack) Push(item interface{}) {
	node := &stackNode{Item: item, Next: nil}
	if stack.Size == 0 {
		stack.head = node
	} else {
		node.Next = stack.head
		stack.head = node
	}
	stack.Size++
}

func (stack *Stack) Pop() (interface{}, bool) {
	var item interface{}
	ok := false
	if !stack.IsEmpty() {
		ok = true
		item = stack.head.Item
		stack.head = stack.head.Next
		stack.Size--
	}
	return item, ok
}

func main() {
	r, output := rand.Perm(50), make([]int, 50)
	stack := NewStack()
	for _, item := range r {
		stack.Push(item)
	}

	fmt.Printf("Size: %v\n", stack.Size)
	fmt.Printf("Items : %v\n\n", r)

	var popped interface{}
	index, ok := 0, true
	for ok == true {
		popped, ok = stack.Pop()
		if ok {
			output[index] = popped.(int)
			index++
		}
	}
	fmt.Printf("Size: %v\n", stack.Size)
	fmt.Printf("Items: %v\n", output)
}
