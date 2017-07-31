package main

import (
	"fmt"
	"math/rand"
)

func Sort(ints []int) {
	if len(ints) > 2 {
		low, middle, high := 0, len(ints)/2, len(ints)-1
		if ints[middle] < ints[low] {
			ints[low], ints[middle] = ints[middle], ints[low]
		}
		if ints[high] < ints[low] {
			ints[high], ints[low] = ints[low], ints[high]
		}
		if ints[high] < ints[middle] {
			ints[high], ints[middle] = ints[middle], ints[high]
		}
		i, j := 0, len(ints)-1
		pivot := ints[j]
		for {
			for ints[i] < pivot {
				i++
			}
			for pivot < ints[j] {
				j--
			}
			if i >= j {
				break
			}
			ints[i], ints[j] = ints[j], ints[i]
		}

		Sort(ints[:i])
		Sort(ints[i:])
	}
}

func main() {
	r := rand.Perm(50)
	Sort(r)
	fmt.Println(r)
}
