package seqsearchst

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type SeqSearchST struct {
	First *Node
}

func NewNode(key string, value interface{}) *Node {
	node := &Node{Key: key, Value: value}
	return node
}

func NewST() *SeqSearchST {
	return &SeqSearchST{}
}

func (st *SeqSearchST) Get(key string) interface{} {
	for i := st.First; i != nil; i = i.Next {
		if i.Key == key {
			return i.Value
		}
	}
	return nil
}

func (st *SeqSearchST) Put(key string, value interface{}) {
	for i := st.First; i != nil; i = i.Next {
		if i.Key == key {
			i.Value = value
		}
	}
	node := NewNode(key, value)
	node.Next, st.First = st.First, node
}
