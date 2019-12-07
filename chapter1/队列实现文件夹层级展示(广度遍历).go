package main

import (
	"data_structure_and_algorithms/chapter1/Queue"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	path:="/home/scorpio/app/golang-SDK/GOPATH/src/MyCode"
	files:=[]string{}
	q:=Queue.NewQueue()
	q.EnQueue(path)
	for !q.IsEmpty(){
		path = q.DeQueue().(string)
		f,err:=ioutil.ReadDir(path)
		if err != nil {
			log.Panic(err)
		}
		for _,v:=range f {
			if v.IsDir() {
				newpath:=path+"/"+v.Name()
				q.EnQueue(newpath)
				files = append(files,newpath )
			}else {
				files = append(files,path+"/"+v.Name() )
			}
		}
	}

	for _,v:=range files {
		fmt.Println(v)
	}
}