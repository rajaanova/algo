package main

type NodeLinkedList struct {
	Node *Node
}
type Node struct {
	data interface{}
	next *Node
}

func NewLinkedList() *NodeLinkedList {
	return &NodeLinkedList{}
}
func (n NodeLinkedList) Add(a int) {
	if n.Node == nil {
		n.Node = &Node{a, nil}
	} else {
		//loop the value
	}
}
func main() {
	list := NewLinkedList()
	list.Add(10)
	list.len()
}
