package main

import (
	"data_structure_and_algorithms/chapter7/mergeMiddleWare"
	"fmt"
	"net"
)

func main() {
	c1 := send("7758")
	c2 := send("7759")
	defer c1.Close()
	defer c2.Close()

	ch1 := reveive(c1)
	ch2 := reveive(c2)
	result := mergeMiddleWare.MergeN(ch1, ch2)
	for v := range result {
		fmt.Println(v)
	}
}

func send(port string) net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		panic(err)
	}
	arry := []byte{byte(0), byte(0)}
	arryR := []byte{byte(0), byte(1)}
	_, err = conn.Write(arry)
	if err != nil {
		panic(err)
	}

	ch := mergeMiddleWare.RandomArrs(50)
	for v := range ch {
		fmt.Println([]byte{byte(1), byte(v)})
		conn.Write([]byte{byte(1), byte(v)})
	}

	_, err = conn.Write(arryR)
	if err != nil {
		panic(err)
	}
	return conn
}

func reveive(conn net.Conn) <-chan int {
	var in chan int
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
			fmt.Println("i1:", i1)
			fmt.Println("i2:", i2)
			if i1 == 0 && i2 == 0 {
				fmt.Println("开始接收！！！")
				in = make(chan int, 1024)
			} else if i1 == 0 && i2 == 1 {
				fmt.Println("接收完成！！！")
				close(in)
				break
			} else if i1 == 1 {
				in <- int(i2)
			}
		}
	}
	return in
}
