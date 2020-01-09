package main

import "fmt"

func main() {
	nums := []int{1, 5, 9, 8, 4, 5, 2, 3, 6, 1, 2, 4, 2, 5, 6, 0, 4, 7, 8, 2, 1, 4, 5, 2, 7, 0, 2, 5, 8, 9, 6, 2, 1, 3, 4, 7, 8, 9, 6, 5, 4, 2, 8, 7, 6, 5, 2, 1, 4, 7, 5, 2, 3, 6}
	shellSort(nums)
	fmt.Println(nums)
}



//真正的希尔排序！
func shellSort(nums []int) {
	length:= len(nums)
	for gap:=length /2 ;gap > 0; gap /= 2 {
		for k:=gap;k < length;k++  {
			for i:= k ; i<length ; i+=gap {
				j := i - gap
				targetNum := nums[i]
				for j >= 0 && nums[j] > targetNum {
					nums[j + gap] ,nums[j] = nums[j] ,nums[j + gap]
					j -= gap
				}
				if targetNum != nums[i] {
					nums[j + gap] = targetNum
				}
			}
		}
	}
}

