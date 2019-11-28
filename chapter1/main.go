package main

import (
	"data_structure_and_algorithms/chapter1/ArrayList"
	"fmt"
)

func main() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	err:=list.Set(0,2)
	if err != nil {
		fmt.Println(err)
	}
	val,err:=list.Get(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	fmt.Println(list)
	err=list.Insert(0,1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
