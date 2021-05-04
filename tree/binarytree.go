/*
@Time : 2021/5/1 下午4:08
@Author : MuYiMing
@File : binarytree
@Software: GoLand
*/
package tree

import "fmt"

type Tree struct {
	Level int
	Root  *Node
}

type Node struct {
	Left  *Node
	Right *Node
	//Data  interface{}
	Val int
}

func NewBinaryTree(val int) *Tree {
	return &Tree{
		Level: 0,
		Root: &Node{
			Val: val,
		},
	}
}

func (this *Node) AddLeftNode(n *Node) *Node {
	this.Left = n
	return this.Left
}

func (this *Node) AddRightNode(n *Node) *Node {
	this.Right = n
	return this.Right
}

func NewBinaryTreeNode(val int, data interface{}) *Node {
	this := Node{}
	this.Val = val
	this.Left = nil
	this.Right = nil
	return &this
}

func (root *Node) MidOrder() {
	if root == nil {
		return
	}

	root.Left.MidOrder()
	fmt.Printf("%d->\n", root.Val)
	root.Right.MidOrder()
}

func (root *Node) PreOrder() {
	if root == nil {
		return
	}
	fmt.Printf("%d->\n", root.Val)
	root.Left.PreOrder()
	root.Right.PreOrder()
}

func (root *Node) PostOrder() {
	if root == nil {
		return
	}
	root.Left.PostOrder()
	root.Right.PostOrder()
	fmt.Printf("%d->\n", root.Val)
}
