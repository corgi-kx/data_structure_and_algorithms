package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 9, 8, 25, 4, 96, 23, 50, 5, 25, 2, 3, 6, 1, 2, 4, 51, 33, 11, 2, 5, 6, 4, 7, 8, 1, 4, 5, 7, 2, 5, 8, 96, 11, 23, 9, 6, 2, 1, 3, 4, 7, 8, 9, 6, 5, 4, 2, 8, 7, 6, 5, 2, 1, 4, 7, 5, 2, 3, 6}
	quickSort(nums)
	fmt.Println(nums)
	fmt.Println(threeFind(nums, 34))
}

func threeFind(nums []int, key int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid1, mid2 := left + (right - left) / 3 , left + (right - left ) / 3 * 2
		if key == nums[mid1] {
			return mid1
		} else if key == nums[mid2] {
			return mid2
		}

		if key < nums[mid1] {
			right = mid1 - 1
		} else if key > nums[mid2] {
			left = mid2 + 1
		} else {
			left = mid1+ 1
			right = mid2 - 1
		}
	}
	return -1
}

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
