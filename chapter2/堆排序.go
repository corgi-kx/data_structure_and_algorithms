package main

import "fmt"

func HeapSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	} else {
		for i := 0; i < length; i++ {
			last := length - i - 1
			HeapSortMax(arr, last)
			arr[0], arr[last] = arr[last], arr[0]
		}
	}
}

//堆排序公式，  左边子节点大小 = 2 × 父节点大小 + 1    右边子节点大小 = 2 × 父节点大小 + 2
func HeapSortMax(arr []int, length int) {
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
	nums := []int{1, 2235, 9, 38, 4, 5, 45512, 3, 622, 1, 22, 44, 22, 5, 1236, 4, 8, 2, 124, 144, 2324, 1235, 2, 73, 5123, 2322, 5553, 1238, 429, 236, 2, 1, 3, 4, 7, 8, 123, 4555, 9, 6, 53, 4, 2, 8, 37, 6, 5, 52, 1, 4, 7, 5, 2, 3, 36}
	HeapSort(nums)
	fmt.Println(nums)
}
