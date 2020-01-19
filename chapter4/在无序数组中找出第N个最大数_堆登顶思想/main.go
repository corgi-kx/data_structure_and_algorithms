package main

import "fmt"


/*
	每次利用二叉树登顶，找出最大值然后剔除，直到找到需要的第N个最大值
*/
func main() {
	nums := []int{95,1,1,9,95,83,10,14,24,36,6,48,57,4,63,8,75,83}
	for i:=1;i<19;i++ {
		fmt.Println(findNmax(i,nums))
	}
}

func findNmax(n int, nums []int) int {
	max := 0
	for i := 1;i<=n;i ++ {
		max = heapMax(nums)
		length:=len(nums)
		nums[0],nums[length -1 ] = nums[length -1],nums[0]
		nums = nums[:length -1]
	}
	return max
}

func heapMax(num []int) (max int) {
	length := len(num)
	for i:=length / 2 ;i>=0;i -- {
		top,left,right:= i,2*i +1,2*i +2
		if left < length && num[left] > num[top] {
			num[top],num[left] = num[left],num[top]
		}
		if right < length && num[right] > num[top] {
			num[top],num[right] = num[right],num[top]
		}
	}
	return num[0]
}