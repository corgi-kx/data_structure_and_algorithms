package main

import (
	"data_structure_and_algorithms/chapter5/link"
	"fmt"
)

func main() {
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
