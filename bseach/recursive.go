/*
@Time : 2021/5/1 下午3:47
@Author : MuYiMing
@File : recursive
@Software: GoLand
*/
package bseach

func BseachRecursive(arr []int, i int) int {
	return helper(arr, 0, len(arr)-1, i)
}

func helper(arr []int, left, right, val int) int {
	if left >= right {
		return -1
	}
	mid := (left + right) / 2

	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		right = mid - 1
	} else {
		left = mid + 1
	}
	return helper(arr, left, right, val)
}
