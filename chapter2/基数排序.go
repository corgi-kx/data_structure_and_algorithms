package main

import "fmt"

func main() {
	nums:=[]int{1,2235,9,38,4,5,45512,3,622,1,22,44,22,5,1236,4,8,2,124,144,2324,1235,2,73,5123,2322,5553,1238,429,236,2,1,3,4,7,8,123,4555,9,6,53,4,2,8,37,6,5,52,1,4,7,5,2,3,36}
	radixSort(nums)
	fmt.Println(nums)
}


func radixSort(nums []int){
	length:= len(nums)
	if length <= 1 {
		return
	}
	maxDigit:=selectMaxDigit(nums)
	bucket := [10][]int{}
	digit:= 1
	for i:=1;i<=maxDigit;i ++ {
		for j:=0;j<length;j++ {
			num:=nums[j] / digit % 10
			bucket[num] = append(bucket[num],nums[j])
		}
		for j,k:=0,0;j<10;j++ {
			bLen :=len(bucket[j])
			if bLen == 0 {
				continue
			}
			for l:=0;l<bLen;l++ {
				nums[k] = bucket[j][l]
				k ++
			}
		}
		bucket = [10][]int{}
		digit = digit * 10
	}
}



//获取最大值的位数
func selectMaxDigit(nums []int) int {
	length:=len(nums)
	//获取最大值
	max:=nums[0]
	for i := 1; i< length; i ++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	//得到最大值的位数
	maxDigest := 0
	for max > 0 {
		maxDigest ++
		max /= 10
	}
	return maxDigest
}
