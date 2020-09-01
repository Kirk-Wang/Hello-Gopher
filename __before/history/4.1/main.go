package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

// (node treeNode)[接受者] 相当于其它语言的 this
// 这里是值传递
func (node treeNode) print() { // print 是用来给这个 node 来接收的
	fmt.Print(node.value, " ")
}

// 编译器会知道你是要值还是指针
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return // 不然会报错
	}
	node.value = value
}

func createNode(value int) *treeNode {
	// 返回的是一个局部变量的地址，C++ 是一个非常典型的错误，但Go不会挂
	return &treeNode{value: value}
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
	// fmt.Println(root) // {0 <nil> <nil>}
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()

	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }

	// fmt.Println(nodes)

	// root.print()
	// root.right.left.setValue(4)
	// root.right.left.print()
	// fmt.Println()

	// root.print()
	// root.setValue(100)

	// pRoot := &root
	// pRoot.print() // 会找到值，取出来一份给它
	// pRoot.setValue(200)
	// pRoot.print()

	// var pRoot *treeNode // nil 指针
	// pRoot.setValue(200)
	// pRoot = &root
	// pRoot.setValue(300)
	// pRoot.print()

}
