package main

import (
	"github.com/Kirk-Wang/Hello-Go/4.2"
)

func main() {
	var root tree.Node
	// fmt.Println(root) // {0 <nil> <nil>}
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
}
