package main

import (
	"data_structure_and_algorithms/chapter1/ArrayList"
	"data_structure_and_algorithms/chapter1/CircleQueue"
	"data_structure_and_algorithms/chapter1/Queue"
	"data_structure_and_algorithms/chapter1/Stack"
	"fmt"
)

func main1() {
	var list ArrayList.Lister = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	list.Append(1)

	fmt.Println(list.String())
	for it:=list.Iterator();it.HasNext();{
		val := it.Next()
		fmt.Println(val)
	}
	fmt.Println(list.String())
}

func main2() {
	var s Stack.Stack = Stack.NewStackList()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	size := s.Size()
	for i:=0;i<size;i++ {
		fmt.Println(s.Pop())
	}
	for i:=0;i<=10000 ; i++ {
		s.Push(i)
	}
	for i:=0;i<size;i++ {
		fmt.Println(s.Pop())
	}
}

func main3() {
	var s ArrayList.Stack = ArrayList.NewStackList()
	s.Push_SK(1)
	s.Push_SK(2)
	s.Push_SK(3)
	s.Push_SK(4)
	size := s.Size_SK()
	for i:=0;i<size;i++ {
		fmt.Println(s.Pop_SK())
	}
	for i:=0;i<=1000 ; i++ {
		s.Push_SK(i)
	}
	for i:=0;i<size;i++ {
		fmt.Println(s.Pop_SK())
	}
}

func main4() {
	var s ArrayList.Stack = ArrayList.NewStackList()
	for i:=0;i<=1000 ; i++ {
		s.Push_SK(i)
	}
	for i:=s.NewIterator();i.HasNext_SK();{
		fmt.Println(i.Next_SK())
	}
}

func main5() {
	fmt.Println(recursion1(10))
}

//用栈来模拟，传入一个数字，递减的值相加
func main6() {
	var s Stack.Stack = Stack.NewStackList()
	s.Push(10)
	data:=0
	for {
		newVal:= s.Pop().(int)
		data += newVal
		if newVal == 0 {
			break
		} else {
			s.Push(newVal - 1)
		}
	}
	fmt.Println(data)
}

func main7() {
	fmt.Println(FBN(12))
}

//利用栈来实现斐波那契数列
func main8() {
	var s Stack.Stack = Stack.NewStackList()
	s.Push(12)
	data:=0
	for !s.IsEmpty(){
		newVal := s.Pop().(int)
		if newVal == 1 || newVal == 2 {
			data += 1
		}else {
			s.Push(newVal - 1 )
			s.Push(newVal - 2)
		}
	}
	fmt.Println(data)
}

//传入一个数字，递减的值相加
func recursion1(num int) int {
	if num == 0 {
		return 0
	}else {
		return num + recursion1(num - 1)
	}
}


//递归实现斐波那契
func FBN(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return  FBN(num - 1 ) + FBN(num - 2)
}


//测试队列
func main9() {
	q:=Queue.NewQueue()
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")
	q.EnQueue("d")
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.IsEmpty())
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")
	q.EnQueue("d")
	fmt.Println(q.Size())
	fmt.Println(q.Front())
	fmt.Println(q.Last())

}

//测试循环队列
func main() {
	cq:=CircleQueue.NewCircleQueue()
	cq.EnQueue("a")
	cq.EnQueue("b")
	cq.EnQueue("c")
	cq.EnQueue("d")
	err:=cq.EnQueue("e")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
}