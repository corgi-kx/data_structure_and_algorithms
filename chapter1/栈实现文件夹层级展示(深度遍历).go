package main

import (
	"data_structure_and_algorithms/chapter1/Stack"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/home/scorpio/app/golang-SDK/GOPATH/src"
	mystack := Stack.NewStackList()
	mystack.Push(path)
	//文件池
	files := make([]string, 0)
	for !mystack.IsEmpty() {
		path = mystack.Pop().(string)
		files = append(files, path)
		fs, _ := ioutil.ReadDir(path)
		for _, f := range fs {
			if f.IsDir() {
				dirName := path + "/" + f.Name()
				mystack.Push(dirName)
			} else {
				fileName := path + "/" + f.Name()
				files = append(files, fileName)
			}
		}
	}

	for _, v := range files {
		fmt.Println(v)
	}
}
