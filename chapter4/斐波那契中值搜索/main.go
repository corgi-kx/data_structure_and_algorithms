package main

import "fmt"

//
//type QQ struct {
//	QQnum      int
//	QQpassword string
//}
//
//var datastoreQQ []QQ
//
//
//func main() {
//	f, err := os.Open(".././data.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	r := bufio.NewReader(f)
//	index :=0
//	for {
//		line, _, err := r.ReadLine()
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//		}
//		id, err2 := strconv.Atoi(strings.Split(string(line), " # ")[0])
//		if err2 != nil {
//			fmt.Println(index)
//			//log.Panic(err2)
//		}
//		password := strings.Split(string(line), " # ")[1]
//		datastoreQQ = append(datastoreQQ, QQ{id, password})
//		index ++
//	}
//	fmt.Println("已将数据加载进切片中...")
//	fmt.Println("开始进行排序...")
//	now := time.Now()
//	quickSort(datastoreQQ)
//	fmt.Println("排序完成,用时：", time.Since(now))
//
//	for {
//		fmt.Println("请输入要搜索的id：")
//		str := ""
//		fmt.Scanln(&str)
//		now := time.Now()
//		searchId, err := strconv.Atoi(str)
//		if err != nil {
//			fmt.Println("输入的不是数字！！！！")
//			continue
//		}
//		fibFind(searchId, datastoreQQ)
//		fmt.Println("用时：", time.Since(now))
//	}
//}
//
//
//func fibFind(id int, data []QQ) {
//	dataLength := len(data)
//	left, right := 0, dataLength-1
//	if id > data[right].QQnum {
//		fmt.Println("传入的数字超过查找范围！！！")
//		return
//	}
//	times:=0
//	for left <= right {
//		times ++
//		leftdata:=float64(id - data[left].QQnum)
//		alldata := float64(data[right].QQnum - data[left].QQnum)
//		gap := float64(right - left)
//		mid := int(float64(left) + leftdata * alldata / gap)
//		if id > data[mid].QQnum {
//			left = mid + 1
//		} else if id < data[mid].QQnum {
//			right = mid - 1
//		} else {
//			fmt.Println("找到了，下标是", mid)
//			fmt.Println(data[mid])
//			fmt.Println("查找次数:",times,"次！")
//			return
//		}
//	}
//	fmt.Println("查找次数:",times,"次！")
//	fmt.Println("没找到！！！")
//}
//
//func quickSort(data []QQ) {
//	length := len(data)
//	if length <= 10 {
//		insertSort(data)
//		return
//	}
//
//	mid, target := 1, data[0]
//	left, right := 0, length-1
//	for left < right {
//		if data[mid].QQnum > target.QQnum {
//			data[mid], data[right] = data[right], data[mid]
//			right--
//		} else {
//			data[mid], data[left] = data[left], data[mid]
//			left++
//			mid++
//		}
//	}
//	quickSort(data[:left])
//	quickSort(data[left+1:])
//}
//
//func insertSort(data []QQ) {
//	length := len(data)
//	for i := 1; i < length; i++ {
//		j := i - 1
//		temp := data[i]
//		for j >= 0 && data[j].QQnum > temp.QQnum {
//			data[j], data[j+1] = data[j+1], data[j]
//			j--
//		}
//		data[j+1], temp = temp, data[j+1]
//	}
//}

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
