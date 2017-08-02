package main

import (
	"fmt"
	"math/rand"
)

type queueNode struct {
	Item interface{}
	Next *queueNode
}

type Queue struct {
	first *queueNode
	last  *queueNode
	Size  int
}

func NewQueue() *Queue {
	return &Queue{first: nil, last: nil, Size: 0}
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size == 0
}

func (queue *Queue) Enqueue(item interface{}) {
	node := &queueNode{Item: item, Next: nil}
	if queue.Size == 0 {
		queue.first = node
		queue.last = node
	} else {
		queue.last.Next = node
		queue.last = node
	}
	queue.Size++
}

func (queue *Queue) Dequeue() (interface{}, bool) {
	var item interface{}
	ok := false
	if !queue.IsEmpty() {
		ok = true
		item = queue.first.Item
		queue.first = queue.first.Next
		queue.Size--
	}
	if queue.IsEmpty() {
		queue.last = nil
	}
	return item, ok
}

func main() {
	r, output := rand.Perm(50), make([]int, 50)
	queue := NewQueue()
	for _, item := range r {
		queue.Enqueue(item)
	}

	fmt.Printf("Size: %v\n", queue.Size)
	fmt.Printf("Items : %v\n\n", r)

	var popped interface{}
	index, ok := 0, true
	for ok == true {
		popped, ok = queue.Dequeue()
		if ok {
			output[index] = popped.(int)
			index++
		}
	}
	fmt.Printf("Size: %v\n", queue.Size)
	fmt.Printf("Items: %v\n", output)
}
