package main

import "fmt"

func main() {
	s:=[]int{1,3,5,10,15,23,45,66,77,78,79,90,100,102,105,107,110,130,135,140,145,153,156,169,175,183,192,200}
	for _,v:=range s {
		fmt.Println(fibFind(s,v))
	}
}

func fib(num int) int {
	if num <=1 {
		return 0
	} else if num == 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

func fibFind(data []int, key int) int {
	dataLength := len(data)
	//fmt.Println(dataLength)
	var fibNum int
	var fibVal int
	for i := 3; i < 20; i++ {
		if dataLength < fib(i) {
			fibNum = i
			fibVal = fib(i)
			break
		}
	}
	//fmt.Println("fibNum", fibNum)
	//fmt.Println("fibVal", fibVal)
	newArr := make([]int, fibVal)
	copy(newArr, data)
	for i := dataLength; i < int(fibVal); i++ {
		newArr[i] = data[dataLength-1]
	}
	//fmt.Println(newArr)
	left, right := 0, dataLength-1
	for left <= right {
		if fibNum < 1 {
			break
		}
		var mid int
		mid = left + fib(fibNum - 1) -1
		//fmt.Println(mid)
		if newArr[mid] < key {
			left = mid + 1
			fibNum -= 2
		} else if newArr[mid] > key {
			right = mid - 1
			fibNum -= 1
		} else {
			return mid
		}
	}
	return -1
}
