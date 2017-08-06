package seqsearchst

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
	max := st.First
	for first := st.First; first != nil; first = first.Next {
		if first.Value.(int) > max.Value.(int) {
			max = first
		}
	}
	fmt.Printf("%v : %v\n", max.Key, max.Value)
}
