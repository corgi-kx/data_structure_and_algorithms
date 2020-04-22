package main

import (
	"fmt"
	"net"
	"sort"
)

func StartReceiving(port string) {
	listen, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var in chan int
	out := make(chan int, 1024)
	for {
		buf := make([]byte, 2)
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				continue
			} else {
				panic(err)
			}
		}

		if n > 0 {
			i1 := buf[0]
			i2 := buf[1]
			if i1 == 0 && i2 == 0 {
				fmt.Println("开始接收！！！")
				in = make(chan int, 1024)
				go IntsSortsToChan(in, out)
			} else if i1 == 0 && i2 == 1 {
				fmt.Println("排序完成！！！")
				close(in)
				break
			} else if i1 == 1 {
				in <- int(i2)
			}
		}
	}
	conn.Write([]byte{0, 0})
	for v := range out {
		fmt.Println(v)
		buf := []byte{1, byte(v)}
		conn.Write(buf)
	}
	conn.Write([]byte{0, 1})
}

func IntsSortsToChan(in, out chan int) {
	nums := []int{}
	for v := range in {
		nums = append(nums, v)
	}
	sort.Ints(nums)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
}

func main() {
	StartReceiving("7758")
}
