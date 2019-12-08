package main

import (
	"fmt"
)

func bubbleSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		isNeedExchange:=false
		for i:=0;i<length -1 ; i ++ {
			for j:=0;j<length - 1 -i;j++ {
				if arr[j] > arr[j+1] {
					arr[j],arr[j+1]=arr[j+1],arr[j]
					isNeedExchange = true
				}
			}
			if !isNeedExchange {
				break
			}
		}
	}
	return arr

}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := bubbleSort(arr)
	fmt.Println(newArr)
}
