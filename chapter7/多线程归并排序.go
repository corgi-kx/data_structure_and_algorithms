package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1:=[]int{1,3,2,5,6,1,5,7,8,2,4,55,66,33,44,66,88,99,22,44,33,55,772,34,6,47,3,5,234,6,8,99,0}
	arr2:=[]int{5,34,4,35,667,8886,234,647,856,452,32667,999,45,8888,44,33,43,535,767,8334,234,777,55}
	sort.Ints(arr1)
	sort.Ints(arr2)

	ch1:=make(chan int,1024)
	ch2:=make(chan int,1024)
	for _,v:=range arr1 {
		ch1 <- v
	}
	for _,v:=range arr2 {
		ch2 <- v
	}

	result:=Merge(ch1,ch2)
	for i,v:=<-result;v ==true;{
		fmt.Println(i)
	}
}

func Merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		in1,v1 :=<-ch1
		in2,v2 :=<-ch2
		for v1 || v2 {
			if !v2 || (v1 && in1 <= in2 ) {
				out <- in1
				in1,v1 =<-ch1
			}else {
				out <- in2
				in2,v2 =<-ch2
			}
		}
	}()
	return out
}
