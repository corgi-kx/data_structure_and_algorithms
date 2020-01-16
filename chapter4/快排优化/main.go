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

type QQ struct {
	QQnum      int
	QQpassword string
}

var datastoreQQ []QQ

func main() {
	f, err := os.Open(".././data.txt")
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
	now := time.Now()
	quickSort(datastoreQQ)
	fmt.Println("排序完成,用时：", time.Since(now))
}

func quickSort(data []QQ) {
	length := len(data)
	if length <= 10 {
		insertSort(data)
		return
	}

	mid, target := 1, data[0]
	left, right := 0, length-1
	for left < right {
		if data[mid].QQnum > target.QQnum {
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

func insertSort(data []QQ) {
	length := len(data)
	for i := 1; i < length; i++ {
		j := i - 1
		temp := data[i]
		for j >= 0 && data[j].QQnum > temp.QQnum {
			data[j], data[j+1] = data[j+1], data[j]
			j--
		}
		data[j+1], temp = temp, data[j+1]
	}
}
