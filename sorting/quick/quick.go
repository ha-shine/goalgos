package quick

func Sort(strs []string) {
	if len(strs) > 2 {
		low, middle, high := 0, len(strs)/2, len(strs)-1
		if strs[middle] < strs[low] {
			strs[low], strs[middle] = strs[middle], strs[low]
		}
		if strs[high] < strs[low] {
			strs[high], strs[low] = strs[low], strs[high]
		}
		if strs[high] < strs[middle] {
			strs[high], strs[middle] = strs[middle], strs[high]
		}
		i, j := 0, len(strs)-1
		pivot := strs[j]
		for {
			for strs[i] < pivot {
				i++
			}
			for pivot < strs[j] {
				j--
			}
			if i >= j {
				break
			}
			strs[i], strs[j] = strs[j], strs[i]
		}

		Sort(strs[:i])
		Sort(strs[i:])
	}
}
