/*
@Time : 2021/5/1 下午4:21
@Author : MuYiMing
@File : bt_test
@Software: GoLand
*/
package tree

import (
	"testing"
)

func TestBtree(t *testing.T) {
	tree := NewBinaryTree(10)
	root := tree.Root
	l1 := NewBinaryTreeNode(5, nil)
	r1 := NewBinaryTreeNode(15, nil)

	l2 := NewBinaryTreeNode(3, nil)
	r2 := NewBinaryTreeNode(7, nil)

	l3 := NewBinaryTreeNode(12, nil)
	r3 := NewBinaryTreeNode(17, nil)

	rootl := root.AddLeftNode(l1)
	rootl.AddLeftNode(l2)
	rootl.AddRightNode(r2)

	rootr := root.AddRightNode(r1)
	rootr.AddLeftNode(l3)
	rootr.AddRightNode(r3)

	tree.Root.MidOrder()
	//fmt.Println("===========")
	//tree.Root.PreOrder()
	//fmt.Println("===========")
	//tree.Root.PostOrder()
	//
	//bs, err := json.Marshal(tree)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%v", string(bs))
	//
	//tmp := &Tree{}
	//err = json.Unmarshal(bs, tmp)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf("%v", tmp.r)
	//tmp.Root.MidOrder()
}
