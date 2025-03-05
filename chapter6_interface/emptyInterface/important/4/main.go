package main

import "fmt"

// 通用类型的节点数据结构

type Node struct {
	data any
	le   *Node
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{nil, left, right}
}
func (n *Node) SetData(data any) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.le = a
	root.ri = b
	fmt.Println(root.data.(string))
	fmt.Println(root.le.data.(string))
	fmt.Println(root.ri.data.(string))
	fmt.Printf("%v\n", root)

}
