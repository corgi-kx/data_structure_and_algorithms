package main

import "fmt"

func main() {
	s := []int{5, 1, 5, 6, 7, 7, 3, 5, 6, 8, 9, 3, 2, 4, 5, 6, 87, 2, 4, 55, 33, 22, 66, 883, 334, 0}
	shellSort(s)

	fmt.Println(s)
}

func shellSort(arr []int) {
	length := len(arr)
	for gap := length / 2; gap > 0; gap /= 2 {
		insertSort(arr, gap)
	}
}

func insertSort(arr []int, gap int) {
	length := len(arr)
	for k := gap; k <= gap*2; k++ {
		go func(k int) {
			for i := 0; i < length-k; i += k {
				j := i + k
				for j > 0 && arr[j] < arr[j-k] {
					arr[j], arr[j-k] = arr[j-k], arr[j]
					j -= k
				}
			}
		}(k)
	}
}
