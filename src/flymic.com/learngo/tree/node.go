package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " + "Ignored.")
	}

	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//func main () {
//	var root Node
//	root = Node{Value: 3}
//	root.Left = &Node {}
//	root.Right = &Node {5, nil, nil}
//	root.Right.Left = new (Node)
//	root.Left.Right = CreateNode(2)
//
//
//	nodes := []Node {
//		{Value: 3},
//		{},
//		{6, nil, &root},
//	}
//	fmt.Println(nodes)
//}
