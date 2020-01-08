package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2235, 9, 38, 4, 5, 45512, 3, 622, 1, 22, 44, 22, 5, 1236, 4, 8, 2, 124, 144, 2324, 1235, 2, 73, 5123, 2322, 5553, 1238, 429, 236, 2, 1, 3, 4, 7, 8, 123, 4555, 9, 6, 53, 4, 2, 8, 37, 6, 5, 52, 1, 4, 7, 5, 2, 3, 36}
	fmt.Println(bucketSort(nums))
}

//桶排序
func bucketSort(nums []int) []int {
	length := len(nums)
	//找出最大值
	max := 0
	for i := 0; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	//入桶
	bucket := make([][]int, length+1)
	for i := 0; i < length; i++ {
		index := nums[i] / max * length //根据公式计算出范围分布
		bucket[index] = append(bucket[index], nums[i])
	}
	//对桶内使用堆排序(或者其他排序都可以)
	nums = []int{}
	for i := 0; i < length+1; i++ {
		if len(bucket[i]) > 0 {
			bucket[i] = heapSort(bucket[i])
			nums = append(nums, bucket[i]...)
		}
	}
	return nums
}

//堆排序
func heapSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		last := length - i - 1
		heapMax(nums, last)
		nums[0], nums[last] = nums[last], nums[0]
	}
	return nums
}

//获取二叉树最大值  nums[0]
func heapMax(nums []int, length int) {
	for i := length / 2; i >= 0; i-- {
		top, left, right := i, 2*i+1, 2*i+2
		if left < length && nums[top] < nums[left] {
			nums[top], nums[left] = nums[left], nums[top]
		}
		if right < length && nums[top] < nums[right] {
			nums[top], nums[right] = nums[right], nums[top]
		}
	}
}
