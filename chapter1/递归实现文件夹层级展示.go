package main

import (
	"fmt"
	"io/ioutil"
)

func addFiles(path string,files []string)  []string{
	files = append(files,path)
	read,_:=ioutil.ReadDir(path)
	for _,f:=range read {
		if f.IsDir() {
			newPath := path+"/"+f.Name()
			files = addFiles(newPath,files)
		}else {
			newPath := path+"/"+f.Name()
			files = append(files,newPath)
		}
	}
	return files
}
func main() {
	path:="/home/scorpio/app/golang-SDK/GOPATH/src"
	files:=[]string{}
	files = addFiles(path,files)
	for _,v:=range files {
		fmt.Println(v)
	}
}
