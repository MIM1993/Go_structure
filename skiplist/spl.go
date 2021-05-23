/*
@Time : 2021/5/22 下午5:12
@Author : MuYiMing
@File : spl
@Software: GoLand
*/
package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//跳表实例化
type SkipList struct {
	head   *Node
	tail   *Node
	size   int64 //跳表总数据量
	levels int64 //跳表总层数
}

//跳表节点
type Node struct {
	id   int64
	val  []byte
	next *Node
	pre  *Node
	up   *Node
	down *Node
}

func NewSkipList() *SkipList {
	spl := &SkipList{}

	spl.head = new(Node)
	spl.head.id = 0

	spl.tail = new(Node)
	spl.tail.id = math.MaxInt64

	spl.head.next = spl.tail
	spl.tail.pre = spl.head

	spl.size = 0
	spl.levels = 1

	return spl
}

//获取跳表数据总量
func (spl *SkipList) Size() int64 {
	return spl.size
}

//获取跳表层数
func (spl *SkipList) Levels() int64 {
	return spl.levels
}

//根据ID获取数据
func (spl *SkipList) Get(id int64) []byte {
	p := spl.findNode(id)
	if p.id == id {
		return p.val
	}
	return nil
}

//插入数据 id重复直接覆盖，暂时不支持相同id
func (spl *SkipList) Insert(id int64, val []byte) {
	p := spl.findNode(id)
	if p.id == id {
		//覆盖
		p.val = val
	}

	//插入节点
	curNode := &Node{
		id:  id,
		val: val,
	}

	spl.insertAfter(p, curNode)

	//通过随机数计算概率，决定是否增加上层索引
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	curLevel := 1

	for rander.Intn(10000) < 5000 {
		curLevel++
		if int64(curLevel) > spl.Levels() {
			spl.newLevel()
		}

		//寻找上层插入位置前节点
		for p.up == nil {
			p = p.pre
		}
		p = p.up

		tmpNode := &Node{
			id: id,
		}

		tmpNode.down = curNode
		curNode.up = tmpNode
		spl.insertAfter(p, tmpNode)

		//改变指针，如果需要继续创建索引层则继续循环
		curNode = tmpNode
	}
	//存储的数据增加
	spl.size++
}

//删除
func (spl *SkipList) Remove(id int64) []byte {
	p := spl.findNode(id)
	if p.id != id {
		return nil
	}

	tmpVal := p.val

	for p != nil {
		p.pre.next = p.next
		p.next.pre = p.pre
		p = p.up
	}

	return tmpVal
}

func (spl *SkipList) findNode(id int64) *Node {
	p := spl.head

	for p != nil {
		if p.id == id {
			//上层表中只有索引，没有数据，去底层取数据\
			if p.down == nil {
				return p
			}
			p = p.down
		} else if p.id < id {
			if p.next.id > id {
				if p.down == nil {
					return p
				}
				p = p.down
			} else {
				p = p.next
			}
		}
	}
	return p
}

func (spl *SkipList) insertAfter(pNode, curNode *Node) {
	curNode.next = pNode.next
	curNode.pre = pNode
	pNode.next.pre = curNode
	pNode.next = curNode
}

func (spl *SkipList) newLevel() {
	head := &Node{
		id: 0,
	}
	tail := &Node{
		id: math.MaxInt64,
	}

	head.next = tail
	tail.pre = head

	head.down = spl.head
	spl.head.up = head

	tail.down = spl.tail
	spl.tail.up = tail

	spl.head = head
	spl.tail = tail
	spl.levels++
}

func (spl *SkipList) print() {
	cacheMap := make(map[int64]int)

	p := spl.head
	for p.down != nil {
		p = p.down
	}

	index := 0
	for p != nil {
		cacheMap[p.id] = index
		p = p.next
		index++
	}

	p = spl.head

	//根据层数打印遍历
	for i := 0; i < int(spl.levels); i++ {
		clNode := p
		preIndex:=0

		for clNode != nil {
			s := clNode.id
			if s== 0{
				fmt.Printf("%s","BRING")
				clNode = clNode.next
				continue
			}

			index := cacheMap[s]
			c := (index - preIndex -1) * 6
			for j:=0;j<c;j++{
				fmt.Print("-")
			}

			if s == math.MaxInt64{
				fmt.Printf("-->%s\n","END")
			}else {
				fmt.Printf("-->%3d",s)
				preIndex = index
			}
			clNode = clNode.next
		}
		p = p.down
	}
}
