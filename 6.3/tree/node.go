package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// (node Node)[接受者] 相当于其它语言的 this
// 这里是值传递
func (node Node) Print() { // print 是用来给这个 node 来接收的
	fmt.Print(node.Value, " ")
}

// 编译器会知道你是要值还是指针
func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return // 不然会报错
	}
	node.Value = Value
}

func CreateNode(Value int) *Node {
	// 返回的是一个局部变量的地址，C++ 是一个非常典型的错误，但Go不会挂
	return &Node{Value: Value}
}
