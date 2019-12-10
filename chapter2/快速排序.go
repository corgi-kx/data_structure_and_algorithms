package main

import "fmt"

func quickSort(values []string) []string{
	if len(values) <= 1 {
		return values
	}
	mid, i := values[0], 1
	left, right := 0, len(values)-1
	for left < right {
		if values[i] > mid {
			values[i], values[right] = values[right], values[i]
			right--
		} else {
			values[i], values[left] = values[left], values[i]
			left++
			i++
		}
	}
	values[left] = mid
	quickSort(values[:left])
	quickSort(values[left+1:])
	return values
}

func main() {
	arr := []string{"i", "a", "c", "z", "n", "f", "q", "b", "t", "j", "f", "d", "p", "s", "e", "z", "h", "a"}
	quickSort(arr)
	fmt.Println(arr)
}
