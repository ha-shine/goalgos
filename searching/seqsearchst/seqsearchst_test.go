package seqsearchst

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestBenchMark(t *testing.T) {
	st := NewST()
	f, _ := os.Open("../tale.txt")
	reader := bufio.NewReader(f)
	word, err := reader.ReadString(' ')
	for err != io.EOF {
		if st.Get(word) != nil {
			st.Put(word, st.Get(word).(int)+1)
		} else {
			st.Put(word, 1)
		}
		word, err = reader.ReadString(' ')
	}
	max := st.First
	for first := st.First; first != nil; first = first.Next {
		if first.Value.(int) > max.Value.(int) {
			max = first
		}
	}
	fmt.Printf("%v : %v", max.Key, max.Value)
}
