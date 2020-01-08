package main

import "fmt"

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	shellSort(arr)
	fmt.Println(arr)
}

func shellSort(nums []string) {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			j := i - gap
			targetNum := nums[i]
			for j >= 0 && nums[j] > targetNum {
				nums[j], nums[j+gap] = nums[j+gap], nums[j]
				j -= gap
			}
			nums[j+gap] = targetNum
		}
	}
}
