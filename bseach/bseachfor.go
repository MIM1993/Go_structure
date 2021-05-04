/*
@Time : 2021/5/1 下午3:43
@Author : MuYiMing
@File : bseachfor
@Software: GoLand
*/
package bseach

func BsearchFor(arr []int, i int) int {
	if len(arr) == 1 {
		return 0
	}

	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	for left <= right {
		if arr[mid] == i {
			return mid
		} else if arr[mid] > i {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return -1
}
