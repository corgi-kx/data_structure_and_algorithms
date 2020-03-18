package main

import "fmt"

func main() {
	s:=[]int{10,2,44,5,667,22,42,0,23,434,546,42,1,23,55,8,996,334,66}
	selectSort(s)
	fmt.Println(s)
}


func bubbleSort(nums []int) {
	length:=len(nums)
	for i:=0;i<length;i++ {
		for j:=0;j<length - i -1 ;j++ {
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}

func selectSort(nums []int) {
	length:=len(nums)
	for i:=0;i<length;i++ {
		target:=i
		for j:=i;j<length -1 ;j++ {
			if nums[target] > nums[j+1] {
				target = j+1
			}
		}
		if target != i {
			nums[i],nums[target] = nums[target],nums[i]
		}
	}
}



func insertSort(nums []int) {
	length:=len(nums)
	for i:=0;i<length -1 ;i++ {
		j:=i+1
		for j>0 && nums[j] < nums[j-1]  {
			nums[j],nums[j-1] = nums[j-1],nums[j]
			j --
		}
	}
}


func mergeSort(nums []int) []int{
	length:=len(nums)
	if length <=1 {
		return nums
	}
	mid:=length/2
	left := mergeSort(nums[:mid])
	right:=mergeSort(nums[mid:])
	return merge(left,right)
}

func merge(left,right []int) []int {
	leftLen,rightLen:=len(left),len(right)
	ln,rn:=0,0
	result:=[]int{}
	for ln< leftLen && rn<rightLen {
		if left[ln] < right[rn] {
			result = append(result,left[ln])
			ln++
		}else {
			result = append(result,right[rn])
			rn++
		}
	}
	if ln<leftLen {
		result = append(result, left[ln:]...)
	}else if  rn<rightLen {
		result = append(result, right[rn:]...)
	}
	return result
}


func quickSort(nums []int) {
	length:=len(nums)
	if length <=1 {
		return
	}
	target,index:=nums[0],1
	left,right:=0,length - 1
	for left<right {
		if nums[index] > target {
			nums[right],nums[index] = nums[index],nums[right]
			right --
		}else {
			nums[left],nums[index] = nums[index],nums[left]
			left ++
			index ++
		}
	}
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}

