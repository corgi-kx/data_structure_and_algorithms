package main

import "fmt"


func main()  {
	nums:=[]int{}
	for i:=300;i<600;i+=1 {
		nums= append(nums, i)
	}
	fmt.Println("斐波那契：")
	fibFind(333,nums)
	fmt.Println("二分法：")
	binarySort(333,nums)
}

func fibFind(id int, data []int) {
	dataLength := len(data)
	left, right := 0, dataLength-1
	if id > data[right] {
		fmt.Println("传入的数字超过查找范围！！！")
		return
	}
	times:=0
	for left <= right {
		times ++
		leftdata:=id - data[left]
		alldata := data[right] - data[left]
		gap := right - left
		mid := left + leftdata * alldata / gap
		if id > data[mid] {
			left = mid + 1
		} else if id < data[mid] {
			right = mid - 1
		} else {
			fmt.Println("找到了，下标是", mid)
			fmt.Println(data[mid])
			fmt.Println("查找次数:",times,"次！")
			return
		}
	}
	fmt.Println("查找次数:",times,"次！")
	fmt.Println("没找到！！！")
}

func binarySort(id int, data []int) {
	dataLength := len(data)
	left, right := 0, dataLength-1
	times:=0
	for left <= right {
		times ++
		mid := (left + right) / 2
		if id > data[mid] {
			left = mid + 1
		} else if id < data[mid] {
			right = mid - 1
		} else {
			fmt.Println("找到了，下标是", mid)
			fmt.Println(data[mid])
			fmt.Println("查找次数：",times,"次！！！")
			return
		}
	}
	fmt.Println("查找次数：",times,"次！！！")
	fmt.Println("没找到！！！")
}
func quickSort(data []int) {
	length := len(data)
	if length <= 10 {
		insertSort(data)
		return
	}

	mid, target := 1, data[0]
	left, right := 0, length-1
	for left < right {
		if data[mid] > target {
			data[mid], data[right] = data[right], data[mid]
			right--
		} else {
			data[mid], data[left] = data[left], data[mid]
			left++
			mid++
		}
	}
	quickSort(data[:left])
	quickSort(data[left+1:])
}

func insertSort(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		j := i - 1
		temp := data[i]
		for j >= 0 && data[j] > temp {
			data[j], data[j+1] = data[j+1], data[j]
			j--
		}
		data[j+1], temp = temp, data[j+1]
	}
}
