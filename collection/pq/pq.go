package main

import (
	"fmt"
	"math/rand"
)

type PQ struct {
	items []int
	Size  int
}

func NewPQ() *PQ {
	return &PQ{items: []int{0}}
}

func (pq *PQ) swim(index int) {
	for index > 1 && pq.items[index/2] < pq.items[index] {
		pq.items[index/2], pq.items[index] = pq.items[index], pq.items[index/2]
		index = index / 2
	}
}

func (pq *PQ) sink(index int) {
	for 2*index <= pq.Size {
		child := index * 2
		if child+1 <= pq.Size && pq.items[child] < pq.items[child+1] {
			child++
		}
		if pq.items[child] < pq.items[index] {
			break
		}
		pq.items[index], pq.items[child] = pq.items[child], pq.items[index]
		index = child
	}
}

func (pq *PQ) Insert(item int) {
	pq.items = append(pq.items, item)
	pq.Size++
	pq.swim(pq.Size)
}

func (pq *PQ) DelMax() int {
	max := pq.items[1]
	pq.items[1], pq.items[pq.Size] = pq.items[pq.Size], pq.items[1]
	pq.Size--
	pq.items = pq.items[:pq.Size+1]
	pq.sink(1)
	return max
}

func main() {
	r := rand.Perm(10)
	pq := NewPQ()
	for _, v := range r {
		pq.Insert(v)
	}
	for pq.Size > 0 {
		fmt.Println(pq.DelMax())
	}
}
