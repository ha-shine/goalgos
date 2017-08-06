package merge

var aux []string

func Sort(a []string) {
	aux = make([]string, len(a))
	sort(a, 0, len(a)-1)
}

func sort(a []string, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, lo, mid)
	sort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func merge(a []string, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		switch {
		case i > mid:
			a[k] = aux[j]
			j++
		case j > hi:
			a[k] = aux[i]
			i++
		case aux[j] < aux[i]:
			a[k] = aux[j]
			j++
		default:
			a[k] = aux[i]
			i++
		}
	}
}
