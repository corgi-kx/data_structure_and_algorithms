package main

import "fmt"

func main() {
	nums := []int{1, 5, 9, 8, 4, 96,23,50,5, 25,2, 3, 6, 1, 2, 4, 51,33,11,2, 5, 6, 4, 7, 8, 2, 1, 4, 5, 2, 7, 2, 5, 8, 9, 6, 2, 1, 3, 4, 7, 8, 9, 6, 5, 4, 2, 8, 7, 6, 5, 2, 1, 4, 7, 5, 2, 3, 6}
	combSort(nums)
	fmt.Println(nums)
}

func combSort(nums []int)  {
	length:=len(nums)
	gap := length
	for gap >= 1 {
		gap = gap * 10 /13
		for i :=0;i+gap < length;i ++ {
			if nums[i] > nums[i+gap] {
				nums[i],nums[i + gap] = nums[i +gap],nums[i]
			}
		}
	}
}
