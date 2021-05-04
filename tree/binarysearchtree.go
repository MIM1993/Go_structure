/*
@Time : 2021/5/2 下午4:59
@Author : MuYiMing
@File : binarysearchtree
@Software: GoLand
*/
package tree

import "fmt"

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func newBSTNode(data int) *BSTNode {
	return &BSTNode{
		data: data,
	}
}

func NewBSTNode(data int) *BSTNode {
	return newBSTNode(data)
}

//插入数据
func (tree *BSTNode) Insert(data int) {
	if tree == nil {
		//tree = newBSTNode(data)
		return
	}
	var p *BSTNode = tree
	for p != nil {
		if p.data > data {
			if p.left == nil {
				p.left = newBSTNode(data)
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = newBSTNode(data)
				return
			}
			p = p.right
		}
	}
}

//中序遍历
func (tree *BSTNode) MidOrder() {
	if tree == nil {
		return
	}
	tree.left.MidOrder()
	fmt.Printf("%d-->", tree.data)
	tree.right.MidOrder()
}

//查找
func (tree *BSTNode) FindNode(val int) *BSTNode {
	if tree == nil {
		return nil
	}
	var p *BSTNode = tree
	for p != nil {
		if p.data == val {
			return p
		} else if p.data > val {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil
}

//删除
func (tree *BSTNode) Delete(val int) {
	if tree == nil {
		return
	}

FLAG:

	var p = tree          // p指向要删除的节点，初始化指向根节点
	var pp = new(BSTNode) // pp记录的是p的父节点

	//查找要删除的节点
	for p != nil && p.data != val {
		pp = p
		if p.data > val {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil { //没找到
		return
	}

	// 要删除的节点有两个子节点
	if p.left != nil && p.right != nil { // 查找右子树中最小节点
		minp := p.right
		minpp := p
		for minp.left != nil {
			minpp = minp
			minp = minp.left
		}
		p.data = minp.data
		p = minp
		pp = minpp
	}

	// 删除节点是叶子节点或者仅有一个子节点
	child := new(BSTNode)
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	//删除节点是根节点,且整个树中只有一个节点
	if pp == nil {
		tree = nil
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

	goto FLAG
}
