package main

import (
	"fmt"
	"reflect"
)

//链表数据类型
type Student struct {
	Id   int
	Age  int
	Addr string
}

//结点结构体
type LinkNode struct {
	//数据域
	Data interface{}
	//指针域
	Next *LinkNode
}

//创建链表
func (node *LinkNode) Create(Data ...interface{}) {
	//容错
	if node == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	//创建头结点
	head := node

	//循环遍历data，取出数据创建链表
	for _, v := range Data {
		//创建新结点
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil

		//将当前结点的下一结点 赋值为新结点
		node.Next = newNode
		//更新新节点的值为当前结点
		node = node.Next
	}
	//将node赋值为头结点
	node = head
}

//打印链表==》递归
func (node *LinkNode) Print1() {
	//递归出口
	if node == nil {
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}

	//递归 打印
	node.Next.Print1()
}

//循环打印链表
func (node *LinkNode) Print2() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next //跳过头结点
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
	}
}

//获取链表长度
func (node *LinkNode) Length() int {
	//容错
	if node == nil {
		return -1
	}
	//定义计数器
	i := 0

	//循环遍历链表
	for node.Next != nil {
		i++
		//后移动
		node = node.Next

	}

	return i
}

//头插法
func (node *LinkNode) InsertByHead(Data interface{}) {
	if node == nil || Data == nil {
		return
	}

	newNode := new(LinkNode)

	newNode.Data = Data
	newNode.Next = nil

	//赋值插入
	newNode.Next = node.Next
	node.Next = newNode

}

//尾插法
func (node *LinkNode) InsertByTail(Data interface{}) {
	//容错
	if node == nil || Data == nil {
		return
	}

	//创建新结点
	newNode := new(LinkNode)
	newNode.Data = Data
	newNode.Next = nil

	//找到尾结点
F:
	for {
		for node.Next == nil {
			node.Next = newNode
			break F
		}
		node = node.Next
	}

	//for node.Next != nil {
	//	node = node.Next
	//}
	//node.Next = newNode
}

//按位置插入
func (node *LinkNode) InsertByIndex(Data interface{}, index int) {
	if node == nil || Data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	if index == node.Length() {
		node.InsertByTail(Data)
		return
	}

	//创建新节点
	newNode := new(LinkNode)
	newNode.Data = Data
	newNode.Next = nil

	//定义变量，记录插入位置的前一个结点
	preNode := node

	//寻找index位置
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}

	//插入
	newNode.Next = node
	preNode.Next = newNode
}

//按位置删除结点
func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}

	//定义前结点
	perNode := node

	for i := 0; i < index; i++ {
		perNode = node
		node = node.Next
	}

	//赋值
	perNode.Next = node.Next

	//置空，驱动GC回收
	node.Data = nil
	node.Next = nil
}

//按数据删除结点
func (node *LinkNode) DeleteByData(Data interface{}) {
	if node == nil || Data == nil {
		return
	}

	//head :=node
	//n:=node.Length()
	////遍历链表
	//for i := 0; i < n; i++ {
	//	node = node.Next
	//	if reflect.DeepEqual(node.Data,Data) && reflect.TypeOf(node.Data) == reflect.TypeOf(Data) {
	//		fmt.Println(i)
	//		head.DeleteByIndex(i+1)
	//		return
	//	}
	//}

	//前一个结点
	preNode := node

	for node.Next != nil {
		preNode = node
		node = node.Next
	}
	//赋值删除
	preNode.Next = node.Next

	//将删除结点置为空
	node.Data = nil
	node.Next = nil
}

//查找数据
func (node *LinkNode) SearchByData(Data interface{}) int {
	if node == nil || Data == nil {
		return -1
	}

	i := 0
	//遍历链表
	for node.Next != nil {
		i++
		node = node.Next
		if reflect.DeepEqual(node.Data, Data) && reflect.TypeOf(node.Data) == reflect.TypeOf(Data) {
			fmt.Println(i)
			return i
		}
	}

	return -1
}

//销毁链表
func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}

	//递归 销毁链表
	node.Next.Destroy()
	node.Next = nil
	node.Data = nil
	node = nil
}

func main() {
	list := new(LinkNode)

	list.Create(1, 2, 3, 4, 5)
	//
	//list.Print1()
	//fmt.Println()
	//list.Print2()
	//fmt.Println()

	//ret := list.Length()
	//fmt.Println(ret)
	//
	//list.InsertByHead(6)
	//list.Print2()

	//list.InsertByTail(9)
	//list.Print1()

	//list.InsertByIndex(10,2)
	//list.Print2()

	//list.DeleteByIndex(0)
	//list.Print2()

	//list.DeleteByData(4)
	//list.Print2()

	//idx := list.SearchByData(5)
	//fmt.Println(idx)

	list.Destroy()
	list.Print2()

}
