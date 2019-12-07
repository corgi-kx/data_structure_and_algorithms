package main

import "fmt"

type node struct {
	data interface{}
	next *node
}

func main() {
	n1:=new(node)
	n1.data = "a"
	n2:=new(node)
	n2.data = "b"
	n2.next = n1
	n3:=new(node)
	n3.data = "c"
	n3.next = n2
	n4:=new(node)
	n4.data = "d"
	n4.next = n3


	fmt.Println(n4.data)
	fmt.Println(n4.next.data)
	fmt.Println(n4.next.next.data)
	fmt.Println(n4.next.next.next.data)
	fmt.Println(n4.next.next.next.next.data)

}