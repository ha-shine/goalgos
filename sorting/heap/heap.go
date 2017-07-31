package main

import (
	"fmt"
	"math/rand"
)

func Sort(ints []int) {
	for i := len(ints) / 2; i >= 0; i-- {
		maxHeapify(ints, i)
	}

	start := len(ints) - 1
	for start >= 0 {
		ints[start], ints[0] = ints[0], ints[start]
		maxHeapify(ints[:start], 0)
		start--
	}
}

func maxHeapify(ints []int, i int) {
	left, right, largest := 2*i+1, 2*i+2, i
	l := len(ints)
	if left < l && ints[left] > ints[largest] {
		largest = left
	}
	if right < l && ints[right] > ints[largest] {
		largest = right
	}

	if largest != i {
		ints[i], ints[largest] = ints[largest], ints[i]
		maxHeapify(ints, largest)
	}
}

func main() {
	ints := rand.Perm(20)
	Sort(ints)
	fmt.Println(ints)
}
