package main

import (
	"fmt"
)

/*
	A.找出某个数字在数组中最小的位置
	B.找出某个数字在数组中最大的位置
	C.找到第一个大于x的数据
	D.找到最后一个小于x的数据
*/
func main() {
	nums := []int{1, 5, 9, 8, 25,4, 96,23,50,5, 25,2, 3, 6, 1, 2, 4, 51,33,11,2, 5, 6, 4, 7, 8, 1, 4, 5, 7, 2, 5, 8,96,11,23, 9, 6, 2, 1, 3, 4, 7, 8, 9, 6, 5, 4, 2, 8, 7, 6, 5, 2, 1, 4, 7, 5, 2, 3, 6}
	quickSort(nums)
	fmt.Println(nums)
	fmt.Println(binarySortA(5,nums))
}

func binarySortA(key int, data []int) int {
	dataLength := len(data)
	left, right := 0, dataLength-1
	times:=0
	for left <= right {
		times ++
		mid := (left + right) / 2
		if key > data[mid] {
			left = mid + 1
		} else if key < data[mid] {
			right = mid - 1
		} else if key == data[mid] && data[mid - 1] == key {
			right = mid - 1
		}else if key == data[mid] && data[mid - 1] != key {
			return mid
		}
	}
	return -1
}

//func binarySortB(key int, data []int) int {
//	dataLength := len(data)
//	left, right := 0, dataLength-1
//	times:=0
//	for left <= right {
//		times ++
//		mid := (left + right) / 2
//		if key > data[mid] {
//			left = mid + 1
//		} else if key < data[mid] {
//			right = mid - 1
//		} else if key == data[mid] && data[mid + 1] == key {
//			left = mid + 1
//		}else if key == data[mid] && data[mid + 1] != key {
//			return mid
//		}
//	}
//	return -1
//}


//func binarySortC(key int, data []int) int {
//	dataLength := len(data)
//	left, right := 0, dataLength-1
//	times:=0
//	for left <= right {
//		times ++
//		mid := (left + right) / 2
//		if mid -1 < 0 {
//			break
//		}
//		if key < data[mid] && key < data[mid - 1] {
//			right = mid -1
//		}else if key >= data[mid]  {
//			left = mid +1
//		} else if key < data[mid] && key >= data[mid -1 ] {
//			return mid
//		}
//	}
//	return -1
//}


//func binarySortD(key int, data []int) int {
//	dataLength := len(data)
//	left, right := 0, dataLength-1
//	times:=0
//	for left <= right {
//		times ++
//		mid := (left + right) / 2
//		if mid +1  >= dataLength {
//			break
//		}
//		if key > data[mid] && key > data[mid + 1 ]  {
//			left = mid + 1
//		}else if key <= data[mid]  {
//			right = mid - 1
//		} else if key > data[mid] &&  key <= data[mid + 1 ]{
//			return mid
//		}
//	}
//	return -1
//}


func quickSort(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}

	mid, target := 1, data[0]
	left, right := 0, length-1

	for left < right {
		if data[mid] > target {
			data[right], data[mid] = data[mid], data[right]
			right--
		} else {
			data[left], data[mid] = data[mid], data[left]
			left++
			mid++
		}
	}
	quickSort(data[:left])
	quickSort(data[left+1:])
}
