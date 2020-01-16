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
	for {
		f, err := os.Open("data.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		fmt.Println("请输入要搜索的id")
		str := ""
		fmt.Scanln(&str)
		target, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("请输入数字id：")
			continue
		}
		now := time.Now()
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			id, _ := strconv.Atoi(strings.Split(string(line), " # ")[0])
			if id == target {
				fmt.Println(string(line))
			}
		}
		fmt.Println("用时：", time.Since(now))
	}
}
