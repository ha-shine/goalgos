package main

import (
	"fmt"
	"math/rand"
)

func Sort(ints []int) {
	m := len(ints) / 2
	if len(ints) > 1 {
		Sort(ints[:m])
		Sort(ints[m:])
		merge(ints)
	}
}

func merge(ints []int) {
	temp := make([]int, len(ints))
	i, l, r := 0, 0, (len(ints) / 2)
	for l < len(ints)/2 && r < len(ints) {
		if ints[l] > ints[r] {
			temp[i] = ints[r]
			r++
		} else if ints[r] > ints[l] {
			temp[i] = ints[l]
			l++
		}
		i++
	}

	for l < len(ints)/2 {
		temp[i] = ints[l]
		i++
		l++
	}

	for r < len(ints) {
		temp[i] = ints[r]
		i++
		r++
	}

	for i := range temp {
		ints[i] = temp[i]
	}
}

func main() {
	r := rand.Perm(50)
	Sort(r)
	fmt.Println(r)
}
