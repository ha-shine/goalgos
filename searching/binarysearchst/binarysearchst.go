package binarysearchst

type BinarySearchST struct {
	keys   []string
	values []interface{}
	Size   int
}

func NewST() *BinarySearchST {
	return &BinarySearchST{Size: 0}
}

func (st BinarySearchST) rank(key string) int {
	lo, hi := 0, st.Size-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case key < st.keys[mid]:
			hi = mid - 1
		case key > st.keys[mid]:
			lo = mid + 1
		default:
			return mid
		}
	}
	return lo
}

func (st BinarySearchST) Get(key string) interface{} {
	if st.Size == 0 {
		return nil
	}
	i := st.rank(key)
	if i < st.Size && st.keys[i] == key {
		return st.values[i]
	}
	return nil
}

func (st *BinarySearchST) Put(key string, value interface{}) {
	i := st.rank(key)
	if i < st.Size && st.keys[i] == key {
		st.values[i] = value
		return
	}
	st.keys = append(st.keys, "None")
	st.values = append(st.values, nil)
	for j := st.Size; j > i; j-- {
		st.keys[j], st.values[j] = st.keys[j-1], st.values[j-1]
	}
	st.keys[i], st.values[i] = key, value
	st.Size++
}

func (st BinarySearchST) Min() string {
	return st.keys[0]
}

func (st BinarySearchST) Max() string {
	return st.keys[st.Size-1]
}

func (st BinarySearchST) Ceiling(key string) string {
	i := st.rank(key)
	return st.keys[i]
}

func (st BinarySearchST) Iterate(lo, hi string, consume func(string)) {
	for i := st.rank(lo); i < st.rank(hi); i++ {
		consume(st.keys[i])
	}
}
