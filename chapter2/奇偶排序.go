package main

import (
	"fmt"
)

//堆排序公式，  左边子节点大小 = 2 × 父节点大小 + 1    右边子节点大小 = 2 × 父节点大小 + 2
func OddEvenSort(arr []string) []string {
	length := len(arr)
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 0; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		for i := 1; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr

}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := OddEvenSort(arr)
	fmt.Println(newArr)
}
