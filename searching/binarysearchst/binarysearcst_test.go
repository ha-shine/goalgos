package binarysearchst

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/ha-shine/goalgos/util"
)

func TestBenchMark(t *testing.T) {
	st := NewST()
	consume := func(s string) {
		rp, _ := regexp.Compile("[a-zA-Z]+")
		s = string(rp.Find([]byte(s)))
		s = strings.ToLower(s)
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

func TestIteration(t *testing.T) {
	st := NewST()
	consume := func(s string) {
		rp, _ := regexp.Compile("[a-zA-Z]+")
		s = string(rp.Find([]byte(s)))
		s = strings.ToLower(s)
		exist := st.Get(s)
		if exist != nil {
			st.Put(s, exist.(int)+1)
		} else {
			st.Put(s, 1)
		}
	}
	util.ReadFile("../../tale.txt", consume)
	printItems := func(item string) {
		fmt.Printf("%v : %v\n", item, st.Get(item))
	}
	st.Iterate(st.keys[0], st.keys[10], printItems)
}
