/*
@Time : 2021/5/22 下午6:40
@Author : MuYiMing
@File : spl_test
@Software: GoLand
*/
package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRander(t *testing.T){
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))


	for i:=0;i<10;i++{
		println(rander.Intn(10000))
	}
}

func TestSpl(t *testing.T){
	spl := NewSkipList()

	spl.Insert(100, []byte("lala"))
	spl.Insert(11, []byte("sx"))
	spl.Insert(266, []byte("11"))
	spl.Insert(3, []byte("dd"))
	spl.Insert(80, []byte("bb"))
	spl.Insert(77, []byte("bb"))
	spl.Insert(65, []byte("bb"))
	spl.Insert(8, []byte("bb"))
	spl.Insert(33, []byte("bb"))
	spl.Insert(47, []byte("bb"))


	spl.print()

	fmt.Println(spl.levels)
	fmt.Println(spl.size)

	data := spl.Get(80)
	fmt.Println(string(data))

	spl.Remove(47)
	spl.print()


}