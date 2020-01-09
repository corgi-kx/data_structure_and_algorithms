package main

import "fmt"

func HeapSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		last := length - i
		for i := last / 2; i >= 0; i-- {
			top, left, right := i, 2*i+1, 2*i+2
			if left < last && nums[top] < nums[left] {
				nums[top], nums[left] = nums[left], nums[top]
			}
			if right < last && nums[top] < nums[right] {
				nums[top], nums[right] = nums[right], nums[top]
			}
		}
		nums[0], nums[last -1] = nums[last -1], nums[0]
	}
}

func main() {
	nums := []int{1, 2235, 9, 38, 4, 5, 45512, 3, 622, 1, 22, 44, 22, 5, 1236, 4, 8, 2, 124, 144, 2324, 1235, 2, 73, 5123, 2322, 5553, 1238, 429, 236, 2, 1, 3, 4, 7, 8, 123, 4555, 9, 6, 53, 4, 2, 8, 37, 6, 5, 52, 1, 4, 7, 5, 2, 3, 36}
	HeapSort(nums)
	fmt.Println(nums)
}
