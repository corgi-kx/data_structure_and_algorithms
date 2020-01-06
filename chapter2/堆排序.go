package main

import (
	"fmt"
)

func HeapSort(arr []string)  {
	length := len(arr)
	if length <= 1 {
		return
	} else {
		for i := 0; i < length; i++ {
			last:=length - i - 1
			HeapSortMax(arr,last)
			arr[0],arr[last] = arr[last],arr[0]
		}
	}
}

//堆排序公式，  左边子节点大小 = 2 × 父节点大小 + 1    右边子节点大小 = 2 × 父节点大小 + 2
func HeapSortMax(arr []string,length int)  {
	//这里需要遍历倒数第二层的根节点,所以需要长度除以二得到的就是最后一个根节点的编号，然后递减遍历所有根节点
	for i := length / 2; i >= 0; i-- {
		topmax := i
		leftchild := 2*topmax + 1
		rightchild := 2*topmax + 2
		//如果左边比我大，我就记录左边
		if leftchild < length && arr[topmax] < arr[leftchild] {
			topmax = leftchild
		}
		//如果右边比我大，我就记录右边
		if rightchild < length && arr[topmax] < arr[rightchild] {
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
}

func main() {
	arr := []string{"i", "a","m","n","e","x","z","r","g","d", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	HeapSort(arr)
	fmt.Println(arr)
}
