package main

import (
	"fmt"
)

func MergeSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		leftArr := MergeSort(arr[:mid])
		rightArr := MergeSort(arr[mid:])
		return merge(leftArr, rightArr)
	}
}

func merge(left, right []string) []string {
	newArr := make([]string,0)
	l,r:=0,0
	for l<len(left) && r < len(right) {
		if left[l] < right[r] {
			newArr = append(newArr,left[l])
			l ++
		}else {
				newArr = append(newArr,right[r])
				r ++
		}
	}
	newArr = append(newArr, left[l:]...)
	newArr = append(newArr, right[r:]...)
	return newArr
}

func main() {
	arr := []string{"r","i", "y","a", "v","m","u","c","k", "n", "f","l", "g","q", "b", "t", "j","x", "d", "p", "s", "e", "z", "w","h", "o"}
	newArr := MergeSort(arr)
	fmt.Println(newArr)
}
