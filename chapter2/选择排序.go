package main

import (
	"fmt"
	"strings"
)

func selectSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if strings.Compare(arr[min], arr[j]) < 0 {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
		return arr
	}
}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := selectSort(arr)
	fmt.Println(newArr)
}
