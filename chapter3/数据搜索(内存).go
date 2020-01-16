package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const threadNum = 10

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
	length := len(datastoreQQ)
	for {
		fmt.Println("请输入要搜索的id：")
		str := ""
		fmt.Scanln(&str)
		id, err2 := strconv.Atoi(str)
		if err2 != nil {
			fmt.Println("请输入数字！")
			continue
		}
		now := time.Now()
		gap := length / threadNum
		index := 0
		wg := sync.WaitGroup{}
		for i := 0; i < threadNum; i++ {
			wg.Add(1)
			go seach(id, datastoreQQ[index:index+gap], &wg)
			index += gap
		}
		wg.Wait()
		fmt.Println("用时：", time.Since(now))
	}
}

func seach(id int, data []QQ, wg *sync.WaitGroup) {
	for i, _ := range data {
		if data[i].QQnum == id {
			fmt.Println("找到了：", data[i])
		}
	}
	wg.Done()
}
