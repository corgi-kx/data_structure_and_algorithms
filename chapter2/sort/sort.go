package main

//快排
func quickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	target, i := nums[0], 1
	left, right := 0, length-1
	for left < right {
		if nums[i] > target {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
		}
	}
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}

//堆排序
func heapSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		last := length - i
		for i := last / 2; i >= 0; i-- {
			top, left, right := i, 2*i+1, 2*i+2
			if left < last && nums[top] < nums[left] {
				nums[top], nums[left] = nums[left], nums[top]
			}
			if right < last && nums[top] < nums[right] {
				nums[top], nums[right] = nums[right], nums[top]
			}
		}
		nums[0], nums[last-1] = nums[last-1], nums[0]
	}
}

//计数排序
func countSort(nums []int) {
	length := len(nums)
	max := 0
	for i := 0; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	bucket := make([][]int, max+1)
	for i := 0; i < length; i++ {
		bucket[nums[i]] = append(bucket[nums[i]], nums[i])
	}
	index := 0
	for i := 0; i < max+1; i++ {
		l := len(bucket[i])
		if l > 0 {
			for j := 0; j < l; j++ {
				nums[index] = bucket[i][j]
				index++
			}
		}
	}
}

//真正的希尔排序！
func shellSort(nums []int) {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for k := gap; k < length; k++ {
			for i := k; i < length; i += gap {
				j := i - gap
				targetNum := nums[i]
				for j >= 0 && nums[j] > targetNum {
					nums[j+gap], nums[j] = nums[j], nums[j+gap]
					j -= gap
				}
				if targetNum != nums[i] {
					nums[j+gap] = targetNum
				}
			}
		}
	}
}

//基数排序
func radixSort(nums []int) {
	length := len(nums)
	max := 0
	for i := 0; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	maxDigit := 0
	for max > 0 {
		maxDigit++
		max /= 10
	}

	digit := 1
	bucket := [10][]int{}
	index := 0
	for digit <= maxDigit {
		for i := 0; i < length; i++ {
			num := nums[i] / digit % 10
			bucket[num] = append(bucket[num], nums[i])
		}
		for i := 0; i < 10; i++ {
			sliceLen := len(bucket[i])
			if sliceLen > 0 {
				for j := 0; j < sliceLen; j++ {
					nums[index] = bucket[i][j]
					index++
				}
			}
		}
		bucket = [10][]int{}
		digit *= 10
	}
}

func insertSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		j := i - 1
		targetNum := nums[i]
		for j >= 0 && nums[j] > targetNum {
			nums[j+1], nums[j] = nums[j], nums[j+1]
			j--
		}
		if targetNum != nums[i] {
			nums[j+1] = targetNum
		}
	}
}

func bucketSort(nums []int) {
	length := len(nums)
	max := 0
	for i := 0; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	bucket := make([][]int, length+1)
	for i := 0; i < length; i++ {
		section := nums[i] / max * length
		bucket[section] = append(bucket[section], nums[i])
	}
	index := 0
	for i := 0; i < length+1; i++ {
		bl := len(bucket[i])
		if bl > 0 {
			quickSort(bucket[i])
			for j := 0; j < bl; j++ {
				nums[index] = bucket[i][j]
				index++
			}
		}
	}
}

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		isExchange := false
		for j := length - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isExchange = true
			}
		}
		if !isExchange {
			break
		}
	}
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	newNums := []int{}
	leftLen, rightLen := len(left), len(right)
	leftIndex, rightIndex := 0, 0
	for leftIndex < leftLen && rightIndex < rightLen {
		if left[leftIndex] < right[rightIndex] {
			newNums = append(newNums, left[leftIndex])
			leftIndex++
		} else {
			newNums = append(newNums, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < leftLen {
		newNums = append(newNums, left[leftIndex:]...)
		leftIndex++
	}
	if rightIndex < rightLen {
		newNums = append(newNums, right[rightIndex:]...)
		rightIndex++
	}
	return newNums
}

func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		targetNum := i
		for j := i + 1; j < length; j++ {
			if nums[targetNum] > nums[j] {
				targetNum = j
			}
		}
		if targetNum != i {
			nums[i], nums[targetNum] = nums[targetNum], nums[i]
		}
	}
}
