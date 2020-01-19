package main

import "fmt"


/*
思路：快速排序。主要思想是找一个“轴”节点，将数列交换变成两部分，
一部分全都小于等于“轴”，另一部分全都大于等于“轴”，然后对两部分 递归处理。其平均时间复杂度是O（NlogN）。
从中可以受到启发，如果我们选择的轴使得交换完的“较大”那一部分的数的个数j正好是n，不也就完成了在 N个数中寻找n个最大的数的任务吗？当然，轴也许不能选得这么恰好。
可以这么分析，如果j>n，则最大的n个数肯定在这j个数中，则问题变成在这j 个数中找出n个最大的数；
否则如果j<n，则这j个数肯定是n个最大的数的一部分，而剩下的j-n个数在小于等于轴的那一部分中，同样可递归处理。这个算法的平均复杂度是O（N）的。
*/
func main() {
	nums := []int{95,1,1,9,95,83,10,14,24,36,6,48,57,4,63,8,75,83}
	for i:=1;i<19;i++ {
		fmt.Println(findNmax(i,nums))
	}
}

func findNmax(n int,nums []int)  int {
	index := grouping(nums)
	if index + 1 == n  {
		return nums[index]
	}else if index + 1> n {
		return findNmax(n,nums[:index])
	}else {
		return findNmax(n - (index +1),nums[index + 1:])
	}
}

//利用快排思想进行分组
func grouping(nums []int) (index int) {
	length:=len(nums)
	if length <= 1 {
		return 0
	}
	mid,target:= 1,nums[0]
	left,right:=0,length - 1
	for left < right {
		if nums[mid] < target {
			nums[mid],nums[right] = nums[right],nums[mid]
			right --
		}else {
			nums[mid],nums[left] = nums[left],nums[mid]
			left ++
			mid ++
		}
	}
	return left
}

