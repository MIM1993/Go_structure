/*
@Time : 2021/5/2 下午5:50
@Author : MuYiMing
@File : bst_test
@Software: GoLand
*/
package tree

import (
	"fmt"
	"testing"
)

func TestBSTNode_Insert(t *testing.T) {
	var tree *BSTNode = NewBSTNode(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)

	t.Run("MidOrder", func(t *testing.T) {
		tree.MidOrder()
	})

	t.Run("find", func(t *testing.T) {
		p := tree.FindNode(99)
		if p == nil {
			t.Error("not find")
			return
		}
		fmt.Println(p.data)
		fmt.Printf("left: %d\n", p.left.data)
		fmt.Printf("right: %d\n", p.right.data)

	})

	t.Run("del", func(t *testing.T) {
		p := tree.FindNode(9)
		fmt.Printf("data: %v\n", p.data)
		tree.Delete(4)
		tree.MidOrder()
	})

	t.Run("repeat", func(t *testing.T) {
		tree.Insert(3)
		tree.MidOrder()
		fmt.Println()
		tree.Delete(3)
		tree.MidOrder()
	})
}
