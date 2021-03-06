package main

import (
	"fmt"
	"goLearning/container/functional/binarytree/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traversal()

	fmt.Println()

	nodeCnt := 0
	root.TraversalFunc(func(n *tree.Node) {
		n.Print()
		nodeCnt++
	})

	fmt.Println("nodeCnt:", nodeCnt)
}
