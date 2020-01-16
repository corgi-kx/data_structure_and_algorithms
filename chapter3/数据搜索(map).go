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

var QQmap = make(map[int]string)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	//最大存入2000万条数据  主要是内存不够用了
	for i := 0; i < 20000000; i++ {
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
		QQmap[id] = password
	}
	fmt.Println("已将数据加载进字典中...")
	for {
		fmt.Println("请输入要搜索的id：")
		str := ""
		fmt.Scanln(&str)
		now := time.Now()
		id, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		if password, ok := QQmap[id]; ok {
			fmt.Println(id, password)
		} else {
			fmt.Println("没找到！")
		}
		fmt.Println("用时：", time.Since(now))
	}
}
