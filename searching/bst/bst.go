package bst

type node struct {
	key         string
	value       interface{}
	left, right *node
	count       int
}

type BST struct {
	root *node
}

func NewNode(key string, value interface{}) *node {
	node := &node{key: key, value: value, left: nil, right: nil, count: 1}
	return node
}

func NewBST() *BST {
	return &BST{root: nil}
}

func (bst *BST) Size() int {
	return size(bst.root)
}

func size(node *node) int {
	if node == nil {
		return 0
	}
	return node.count
}

func (bst *BST) Get(key string) interface{} {
	current := bst.root
	for current != nil {
		switch {
		case key < current.key:
			current = current.left
		case key > current.key:
			current = current.right
		default:
			return current.value
		}
	}
	return nil
}

func (bst *BST) Put(key string, value interface{}) {
	bst.root = put(bst.root, key, value)
}

func put(node *node, key string, value interface{}) *node {
	if node == nil {
		return NewNode(key, value)
	}
	if key < node.key {
		node.left = put(node.left, key, value)
	} else if key > node.key {
		node.right = put(node.right, key, value)
	} else {
		node.value = value
	}
	node.count = size(node.left) + size(node.right) + 1
	return node
}
