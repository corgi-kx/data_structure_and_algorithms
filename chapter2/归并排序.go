package main

import "fmt"

func mergeSort(nums []int) []int{
	length:=len(nums)
	if length <= 1 {
		return nums
	}
	mid := length /2
	left:= mergeSort(nums[mid:])
	right := mergeSort(nums[:mid])
	return merge(left,right)
}

func merge(left,right []int) (newNum []int) {
	li,ri:=0,0
	ll,rl:=len(left),len(right)
	for li <ll && ri <rl {
		if left[li] < right[ri] {
			newNum = append(newNum,left[li])
			li ++
			continue
		}
		newNum = append(newNum,right[ri])
		ri ++
	}
	if li <ll {
		newNum = append(newNum,left[li:]...)
	}else if ri < rl {
		newNum = append(newNum,right[ri:]...)
	}
	return
}

func main() {
	nums:=[]int{1,5,9,8,4,5,2,3,6,1,2,4,2,5,6,4,7,8,2,1,4,5,2,7,2,5,8,9,6,2,1,3,4,7,8,9,6,5,4,2,8,7,6,5,2,1,4,7,5,2,3,6}
	fmt.Println(mergeSort(nums))
}