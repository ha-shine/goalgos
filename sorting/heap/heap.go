package heap

func Sort(strs []string) {
	for i := len(strs) / 2; i >= 0; i-- {
		maxHeapify(strs, i)
	}

	start := len(strs) - 1
	for start >= 0 {
		strs[start], strs[0] = strs[0], strs[start]
		maxHeapify(strs[:start], 0)
		start--
	}
}

func maxHeapify(strs []string, i int) {
	left, right, largest := 2*i+1, 2*i+2, i
	l := len(strs)
	if left < l && strs[left] > strs[largest] {
		largest = left
	}
	if right < l && strs[right] > strs[largest] {
		largest = right
	}

	if largest != i {
		strs[i], strs[largest] = strs[largest], strs[i]
		maxHeapify(strs, largest)
	}
}
