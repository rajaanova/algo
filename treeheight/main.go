package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {

	root := &Node{left: &Node{data: 2}, right: &Node{data: 7}, data: 3}
	root.left.left = &Node{data: 1}
	root.left.right = &Node{data: 5}
	i := getSum(root, 0)
	fmt.Println(i)
}
func getSum(parent *Node, partial_sum int) int {
	if parent == nil {
		return partial_sum
	}
	partial_sum = partial_sum + parent.data
	if parent.left != nil && parent.right != nil {
		partial_sum = partial_sum + parent.data
		partial_sum = getSum(parent.left, partial_sum)
		partial_sum = getSum(parent.right, partial_sum)
		return partial_sum
	} else if parent.left != nil {
		partial_sum = getSum(parent.left, partial_sum)
	} else {
		partial_sum = getSum(parent.right, partial_sum)
	}
	return partial_sum
}
func getHeight(parent *Node, data int) int {
	if parent == nil {
		return -1
	}
	if parent.data == data {
		return 0
	}
	if x := getHeight(parent.left, data); x >= 0 {
		return x + 1
	}
	if x := getHeight(parent.right, data); x > 0 {
		return x + 1
	}
	return -1
}
