/*
@Time : 2021/5/1 下午3:43
@Author : MuYiMing
@File : bseach_test
@Software: GoLand
*/
package bseach

import (
	"fmt"
	"testing"
)

func TestBsearchFor(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 5
	idx := BsearchFor(arr, i)
	fmt.Println(idx)
	fmt.Println(arr[idx])
}

func TestBseachRecursive(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 5
	idx := BseachRecursive(arr, i)
	fmt.Println(idx)
	fmt.Println(arr[idx])
}
