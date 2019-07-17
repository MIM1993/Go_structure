package main

import (
	"fmt"
)

//双向链表结点
type LinkDoubleNode struct {
	Data interface{}
	Prev *LinkDoubleNode
	Next *LinkDoubleNode
}

//创建双向链表
func (node *LinkDoubleNode) Create(Data ...interface{}) {
	if node == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	//保存头结点
	head := node

	//循环Data
	for _, v := range Data {
		newNode := new(LinkDoubleNode)
		newNode.Data = v
		newNode.Prev = node
		newNode.Next = nil

		//当前结点的下一个结点赋值为新节点
		node.Next = newNode
		//将当前结点赋值为下一个结点
		node = node.Next
	}
	//还原头结点
	node = head
}

//打印双向链表 --->正向递归
func (node *LinkDoubleNode) Print1() {
	if node == nil { //容错  递归出口
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}

	node.Next.Print1()
}

//打印双向链表 --->逆向循环打印
func (node *LinkDoubleNode) Print2() {
	if node == nil {
		return
	}

	//找到尾节点
	for node.Next != nil {
		node = node.Next
	}

	//循环 prev 倒叙 打印
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		//node前移
		node = node.Prev
	}
}

//获取长度
func (node *LinkDoubleNode) Length() int {
	if node == nil {
		return -1
	}

	//定义计数器
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

//插入==》按位置插入
func (node *LinkDoubleNode) InsertByIndex(Data interface{}, index int) {
	if node == nil || Data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}

	//定义index前结点
	perNode := node
	//循环找到index位置的node
	for i := 0; i < index; i++ {
		perNode = node
		node = node.Next
	}

	//定义新结点
	newNode := new(LinkDoubleNode)
	newNode.Data = Data

	//正序添加
	perNode.Next = newNode
	newNode.Next = node

	//反序添加
	node.Prev = newNode
	newNode.Prev = perNode
}

//删除结点==>按index删除
func (node *LinkDoubleNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}

	//定义变量，存储链表长度
	l := node.Length()

	//定义前驱结点
	preNode := node

	//寻找位置结点
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}

	//如果index是最后一个结点
	if index == l {
		preNode.Next = nil
		node.Data = nil
		node.Prev = nil
		node.Next = nil
		node = nil
	} else {
		//赋值，删除结点
		preNode.Next = node.Next
		node.Next.Prev = preNode

		//置空，结点
		node.Data = nil
		node.Prev = nil
		node.Next = nil
		node = nil
	}
}

//销毁链表
func (node *LinkDoubleNode) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Next = nil
	node.Prev = nil
	node.Data = nil
	node = nil
}

func main() {
	list := new(LinkDoubleNode)
	list.Create(1, 2, 3, 4, 5)

	//list.Print1()
	//fmt.Println()
	//list.Print2()

	//len := list.Length()
	//fmt.Println(len)

	//list.InsertByIndex(777, 6)
	list.DeleteByIndex(5)
	list.Print1()
	fmt.Println()
	list.Print2()
	list.Destroy()
	list.Print1()
	list.Print2()
}
