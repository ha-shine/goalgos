package main

import (
	"fmt"
	"math/rand"
)

type Stack struct {
	Items  []interface{}
	Size   int
	Length int
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.Items = make([]interface{}, 5)
	stack.Size = 5
	stack.Length = 0
	return stack
}

func (stack *Stack) resize(size int) {
	newItems := make([]interface{}, size)
	for i := 0; i < len(newItems) && i < len(stack.Items); i++ {
		newItems[i] = stack.Items[i]
	}
	stack.Items = newItems
	stack.Size = size
}

func (stack *Stack) grow() {
	stack.resize(stack.Size * 2)
}

func (stack *Stack) shrink() {
	stack.resize(stack.Size / 2)
}

func (stack *Stack) Push(item interface{}) {
	if stack.Length == stack.Size {
		stack.grow()
	}
	stack.Items[stack.Length] = item
	stack.Length++
}

func (stack *Stack) Pop() (interface{}, bool) {
	if stack.Length != 0 {
		item := stack.Items[stack.Length-1]
		stack.Length--
		if stack.Size/2 >= stack.Length {
			stack.shrink()
		}
		return item, true
	}
	return nil, false
}

func main() {
	r := rand.Perm(50)
	stack := NewStack()
	for _, item := range r {
		stack.Push(item)
	}
	fmt.Printf("Size: %v, Length: %v\n%v\n\n", stack.Size, stack.Length, stack.Items)
	ok := true
	for ok == true {
		_, ok = stack.Pop()
	}
	fmt.Printf("Size: %v, Length: %v\n%v\n", stack.Size, stack.Length, stack.Items)
}
