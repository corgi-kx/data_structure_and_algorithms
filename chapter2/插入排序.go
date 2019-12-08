package main

import (
	"fmt"
	"strings"
)

func insertSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup:=arr[i]
			j:=i - 1
			for j>=0 && strings.Compare(backup,arr[j]) < 0   {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
	}
	return arr

}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := insertSort(arr)
	fmt.Println(newArr)
}
