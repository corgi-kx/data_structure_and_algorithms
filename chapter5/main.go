package main

import (
	"data_structure_and_algorithms/chapter5/doubleLink"
	"data_structure_and_algorithms/chapter5/link"
	"data_structure_and_algorithms/chapter5/priorityQueue"
	"fmt"
)

func main1() {
	list:=link.NewLinkList()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add("f")
	list.Add("j")
	list.Add("h")
	list.Add("i")
	list.Add("g")
	list.Add("k")
	list.Add("l")
	list.Add("m")
	list.Add("n")
	list.Add("o")

	fmt.Println(list.String())
	//fmt.Println(list.GetFirstNode())
	//fmt.Println(list.GetRearNode())
	//list.InsertFront("haha")
	//list.InsertRear("heihei")
	//fmt.Println(list.String())
	//list.InsertIndex(5,"呵呵呵")
	//list.InsertRear("嘎嘎嘎")
	//fmt.Println(list.String())
	//fmt.Println(list.GetMidNode())
	//node:=link.NewLinkNode("C")
	//list.ReplayNode(list.GetMidNode(),node)
	//fmt.Println(list.String())
	//list.UpdateByIndex(2,"B")
	//fmt.Println(list.String())
	//fmt.Println(list.GetSize())
	//list.DeleteByNode(list.GetFirstNode())
	//fmt.Println(list.String())
	//list.DeleteByIndex(2)
	//fmt.Println(list.String())
	//fmt.Println(list.GetSize())
	list.Reversal()
	fmt.Println(list.String())
}

func main2() {
	l:= doubleLink.NewDoubleLinkList()
	l.Add("haha")
	l.Add("heihei")
	l.Add("hehe")
	l.Add("houhou")
	fmt.Println(l.String())
	fmt.Println(l.GetSize())
	l.InsertFront("gaga")
	fmt.Println(l.String())
	fmt.Println(l.GetSize())
	l.InsertFront("what")
	fmt.Println(l.String())
	fmt.Println(l.GetSize())
	l.InsertIndex(0,"nihao")
	fmt.Println(l.String())
	fmt.Println(l.GetSize())
	l.InsertIndex(6,"oooo")
	fmt.Println(l.String())
	fmt.Println(l.GetSize())
	fmt.Println(l.ReversePrint())
}

func main() {
	q:=priorityQueue.New()
	q.Enqueue(11,"中国")
	q.Enqueue(33,"日本")
	q.Enqueue(55,"韩国")
	q.Enqueue(59,"朝鲜")
	q.Enqueue(644,"美国")
	q.Enqueue(763,"泰国")
	q.Enqueue(0,"古巴")
	q.Enqueue(8,"希腊")
	q.Enqueue(12,"法国")
	q.Enqueue(76,"英国")
	q.Enqueue(86,"印度")
	q.Enqueue(94,"利比亚")
	q.Enqueue(56,"加拿大")
	q.Enqueue(76,"英国2")
	q.Enqueue(56,"加拿大2")
	q.Enqueue(8,"希腊2")
	q.Enqueue(763,"泰国2")


	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}