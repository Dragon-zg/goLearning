package tree

import "fmt"

//定义二叉树
type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil, ignore!")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//中序遍历二叉树
func (node *Node) Traversal() {
	if node == nil {
		return
	}
	node.Left.Traversal()
	node.Print()
	node.Right.Traversal()
}

//中序遍历二叉树
func (node *Node) TraversalFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversalFunc(f)
	f(node)
	node.Right.TraversalFunc(f)
}
