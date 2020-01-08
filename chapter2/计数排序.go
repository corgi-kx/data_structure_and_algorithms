package main

import "fmt"

func main() {
	nums := []int{2235, 9, 38, 4, 5, 45512, 3, 622, 1, 22, 44, 22, 5, 1236, 4, 8, 2, 124, 144, 2324, 1235, 2, 73, 5123, 2322, 5553, 1238, 429, 236, 2, 1, 3, 4, 7, 8, 123, 4555, 9, 6, 53, 4, 2, 8, 37, 6, 5, 52, 1, 4, 7, 5, 2, 3, 36}
	countSort(nums)
	fmt.Println(nums)
}

func selectMaxAndMin(nums []int) (max, min int) {
	length := len(nums)
	max, min = nums[0], nums[0]
	for i := 0; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		} else if min > nums[i] {
			min = nums[i]
		}
	}
	return
}

func countSort(nums []int) {
	length := len(nums)
	max, _ := selectMaxAndMin(nums)
	arr := make([]int, max+1)
	for i := 0; i < length; i++ {
		arr[nums[i]]++
	}
	index := 0
	for i := 0; i < max+1; i++ {
		for arr[i] != 0 {
			nums[index] = i
			index++
			arr[i]--
		}
	}
}

//
//func countSort(arr []int) {
//	maxVal := 0
//	length := len(arr)
//	for i := 0; i < length; i ++ {
//		if arr[i] > maxVal {
//			maxVal = arr[i]
//		}
//	}
//	tmp := make([]int, maxVal+1)
//	for i := 0; i < length; i ++ {
//		tmp[arr[i]] += 1
//	}
//	j := 0
//	for i := 0; i < maxVal+1; i ++ {
//		for tmp[i] > 0 {
//			arr[j] = i
//			j ++
//			tmp[i] --
//		}
//	}
//}
