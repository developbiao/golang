package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"

	"flymic.com/learngo/tree"
)

// extends class
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	root.Traverse()

	fmt.Println("My own post-order tranversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count: ", nodeCount)

	testPrint()
	fmt.Println("ok")
}

func testPrint() {
	log.Print(log.WARNING, "%s Hello test Log")
}
