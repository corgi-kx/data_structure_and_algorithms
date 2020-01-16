package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		id, err2 := strconv.Atoi(strings.Split(string(line), " # ")[0])
		if err2 != nil {
			panic(err2)
		}
		password := strings.Split(string(line), " # ")[1]
		datastoreQQ = append(datastoreQQ, QQ{id, password})
	}
	fmt.Println("已将数据加载进切片中...")
	fmt.Println("开始进行排序...")
	quickSort(datastoreQQ)
	fmt.Println("排序完成")

	for {
		fmt.Println("请输入要搜索的id：")
		str := ""
		fmt.Scanln(&str)
		now := time.Now()
		searchId, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("输入的不是数字！！！！")
			continue
		}
		binarySort(searchId, datastoreQQ)
		fmt.Println("用时：", time.Since(now))
	}
}

func binarySort(id int, data []QQ) {
	dataLength := len(data)
	left, right := 0, dataLength-1
	times:=0
	for left <= right {
		times ++
		mid := (left + right) / 2
		if id > data[mid].QQnum {
			left = mid + 1
		} else if id < data[mid].QQnum {
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

func quickSort(data []QQ) {
	length := len(data)
	if length <= 1 {
		return
	}

	mid, target := 1, data[0]
	left, right := 0, length-1

	for left < right {
		if data[mid].QQnum > target.QQnum {
			data[right], data[mid] = data[mid], data[right]
			right--
		} else {
			data[left], data[mid] = data[mid], data[left]
			left++
			mid++
		}
	}
	quickSort(data[:left])
	quickSort(data[left+1:])
}
