package main

import (
	"data_structure_and_algorithms/chapter1/ArrayList"
	"fmt"
)

func main() {
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
