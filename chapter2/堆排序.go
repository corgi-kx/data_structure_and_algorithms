package main

import (
	"fmt"
)

func HeapSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			arr[i] = HeapSortMax(arr[i+1:])
		}
	}
	return arr
}

//堆排序公式，  左边子节点大小 = 2 × 父节点大小 + 1    右边子节点大小 = 2 × 父节点大小 + 2
func HeapSortMax(arr []string) string {
	length := len(arr)
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
	return arr[0]

}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	newArr := HeapSort(arr)
	fmt.Println(newArr)
}
