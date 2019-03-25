package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	fmt.Println(root)
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root)

	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()
	root.traverse()
}
