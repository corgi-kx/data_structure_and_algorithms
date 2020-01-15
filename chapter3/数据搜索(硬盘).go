package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		f, err := os.Open("data.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		fmt.Println("请输入要搜索的字符串：")
		str := ""
		fmt.Scanln(&str)
		now:=time.Now()
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			pass := strings.Split(string(line), " # ")[1]
			if strings.Contains(pass, str) {
				fmt.Println(pass)
			}
		}
		fmt.Println("用时：",time.Since(now))
	}
}
