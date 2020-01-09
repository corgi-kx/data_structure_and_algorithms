package main

import (
	"crypto/rand"
	"math/big"
	"testing"
	"time"
)



var size int64 =100000
func TestSortTimes(t *testing.T) {
	t.Log("====================将十万随机数进行排序====================")
	nums:=randNums()
	sortST := time.Now()
	bubbleSort(nums)
	sortET := time.Since(sortST)
	t.Logf("===================冒泡排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	selectSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================选择排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	insertSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================插入排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	shellSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================希尔排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	quickSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================快速排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	heapSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================堆排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	mergeSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================归并排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	radixSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================基数排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	countSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================计数排序用时:%fs===================",sortET.Seconds())
	nums =randNums()
	sortST = time.Now()
	bucketSort(nums)
	sortET = time.Since(sortST)
	t.Logf("===================桶排序用时:%fs===================",sortET.Seconds())
}




func randNums() []int{
	nums:= make([]int,size)
	bSize:=big.NewInt(size)
	for i:=0;i<len(nums);i++ {
		bInt,_ := rand.Int(rand.Reader,bSize)
		nums[i] =int(bInt.Int64())
	}
	return nums
}