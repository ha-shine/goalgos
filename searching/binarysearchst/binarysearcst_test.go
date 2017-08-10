package binarysearchst

import (
	"fmt"
	"testing"

	"github.com/ha-shine/goalgos/util"
)

func TestBenchMark(t *testing.T) {
	st := NewST()
	consume := func(s string) {
		exist := st.Get(s)
		if exist != nil {
			st.Put(s, exist.(int)+1)
		} else {
			st.Put(s, 1)
		}
	}
	util.ReadFile("../../tale.txt", consume)
	max := 0
	for i := 0; i < st.Size; i++ {
		if st.values[max].(int) < st.values[i].(int) {
			max = i
		}
	}
	fmt.Printf("%v : %v\n", st.keys[max], st.values[max])
}
