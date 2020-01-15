package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var datastore []string
const threadNum  = 10


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
		pass := strings.Split(string(line), " # ")[1]
		datastore = append(datastore,pass)
	}
	fmt.Println("已将数据加载进切片中...")
	length:=len(datastore)
	for {
		fmt.Println("请输入要搜索的字符串：")
		str := ""
		fmt.Scanln(&str)
		now:=time.Now()
		gap := length / threadNum
		index:=0
		wg:=sync.WaitGroup{}
		for i:=0;i<threadNum;i++ {
			wg.Add(1)
			go seach(str,datastore[index:index+gap],&wg)
			index +=gap
		}
		wg.Wait()
		fmt.Println("用时：",time.Since(now))
	}
}

func seach(target string,data []string,wg *sync.WaitGroup) {
	for i,_:=range data {
		if strings.Contains(data[i],target) {
			fmt.Println(data[i])
		}
	}
	wg.Done()
}