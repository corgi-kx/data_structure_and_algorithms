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

}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := OddEvenSort(arr)
	fmt.Println(newArr)
}
