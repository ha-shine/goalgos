package merge

func Sort(strs []string) {
	m := len(strs) / 2
	if len(strs) > 1 {
		Sort(strs[:m])
		Sort(strs[m:])
		merge(strs)
	}
}

func merge(strs []string) {
	temp := make([]string, len(strs))
	i, l, r := 0, 0, (len(strs) / 2)
	for l < len(strs)/2 && r < len(strs) {
		if strs[l] > strs[r] {
			temp[i] = strs[r]
			r++
		} else if strs[r] > strs[l] {
			temp[i] = strs[l]
			l++
		}
		i++
	}

	for l < len(strs)/2 {
		temp[i] = strs[l]
		i++
		l++
	}

	for r < len(strs) {
		temp[i] = strs[r]
		i++
		r++
	}

	for i := range temp {
		strs[i] = temp[i]
	}
}
