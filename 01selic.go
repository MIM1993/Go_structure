package main

/*
#include <stdlib.h>
*/
import "C" //导入Ｃ语言库

import (
	"unsafe"
	"fmt"
)

//定义切片类型
type Slice struct {
	Data unsafe.Pointer //GO 语言中的万能指针  没有具体数据类型，不能进行运算
	Len  int            //数据元素的个数
	Cap  int            //可扩展的有效容量
}

//一个字符的位数
const TAG = 8

//创建　　切片
func (s *Slice) Create(l int, c int, Data ...int) {
	//容错
	if s == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}
	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}

	//申请内存空间 Ｃ＝＝》函数  一个int 类型　８　个字节  c语言返回viod*　类型,不能运算
	s.Data = C.malloc(C.size_t(c) * 8)

	//添加数据
	s.Len = l
	s.Cap = c

	//首先　将　s.Data（地址值，不能运算）　转换为可以计算的数值（大数）
	p := uintptr(s.Data)

	//根据Data　遍历　存入内存中
	for _, v := range Data {
		//必须是p进行转换
		//将p(数值)转化为指针(地址)，再强转为int指针类型，最后解引用，
		// 获取内存，然后赋值
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}

	//释放内存
	//C.free(s.Data)
}

//打印切片
func (s *Slice) Print() {
	//容错
	if s == nil {
		return
	}

	//首先　将　s.Data（地址值，不能运算）　转换为可以计算的数值（大数）
	p := uintptr(s.Data)
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += TAG
	}
}

//扩展容量　追加元素　
func (s *Slice) Append(Data ...int) {
	//容错
	if s == nil {
		return
	}
	//循环判断是否需要扩容,扩展为两倍
	for s.Len+len(Data) > s.Cap {
		//扩容函数　　必须类型强转　
		C.realloc(s.Data, C.size_t(s.Cap)*2*TAG) //记得乘８
		s.Cap *= 2
	}

	//转换
	p := uintptr(s.Data)

	//偏移
	p += uintptr(s.Len * TAG)

	//循环添加
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}

	//修改额len长度
	s.Len += len(Data)
}

//获取元素
func (s *Slice) GetData(index int) int {
	//容错
	if s == nil {
		return -1
	}
	if index < 0 && index >= s.Len {
		return -1
	}

	//获取元素
	//类型转换
	p := uintptr(s.Data)

	//偏移
	p += uintptr(index * TAG)

	//取数据
	return *(*int)(unsafe.Pointer(p))
}

//已知元素，返回下标
func (s *Slice) SearchData(Data int) int {
	if s == nil {
		return -1
	}

	//转换
	p := uintptr(s.Data)

	//循环
	for i := 0; i < s.Len; i++ {
		if *(*int)(unsafe.Pointer(p)) == Data {
			//返回下标
			return i
		}
		p += uintptr(TAG)
	}

	//没找到
	return -1
}

//删除元素
func (s *Slice) DeleteData(index int) bool {
	//容错
	if s == nil {
		return false
	}
	if index < 0 && index >= s.Len {
		return false
	}

	//确定下标指针
	p := uintptr(s.Data)
	p += uintptr(index) * TAG

	//后一个指针
	after := p

	//循环赋值
	for i := index; i < s.Len; i++ {
		after += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(after))
		p += TAG
	}

	//修改len
	s.Len -= 1

	return true
}

//根据下标插入数据
func (s *Slice) Insert(Data int, index int) bool {
	//容错
	if s == nil {
		return false
	}
	if index < 0 && index > s.Len {
		return false
	}
	//如果插入位置在结尾，直接追加
	if index == s.Len {
		s.Append(Data)
		return true
	}
	//循环判断是否需要扩容,扩展为两倍
	if s.Len == s.Cap {
		//扩容函数　　必须类型强转　
		C.realloc(s.Data, C.size_t(s.Cap)*2*TAG) //记得乘８
		s.Cap *= 2
	}

	//插入位置在中间
	p := uintptr(s.Data)

	//将ｐ偏移
	p += TAG * uintptr(index)

	//获取完成后的最后一个元素
	temp := p + TAG*uintptr(s.Len)

	//将index之后的元素向后移动
	for i := s.Len; i > index; i-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}

	//插入元素
	*(*int)(unsafe.Pointer(p)) = Data

	//修改len
	s.Len++

	return true
}

//销毁切片
func (s *Slice) Destroy() {
	//当一个元素为nil时，垃圾回收机制会进行内存回收
	if s == nil || s.Data == nil {
		return
	}

	//释放内存
	C.free(s.Data)

	//设置为nil,驱动go语言进行垃圾回收
	s.Data = nil
	s.Len = 0
	s.Cap = 0

	//手动回收
	//runtime.GC()
}

func main() {
	slice := new(Slice)
	slice.Create(6, 10, 1, 2, 3, 4, 5, 6)
	slice.Append(11, 22, 33)
	//slice.Print()
	//根据下标获取数据
	//ret :=slice.GetData(5)
	//ret := slice.SearchData(1)

	//ret := slice.DeleteData(1)

	ret := slice.Insert(999, 0)

	fmt.Println(ret)

	slice.Print()

	fmt.Println("----------------------")
	slice.Destroy()
	slice.Print()

}
