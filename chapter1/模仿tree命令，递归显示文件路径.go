package main

import (
	"fmt"
	"io/ioutil"
)

func addFilesX(path string, files []string, level int) []string {
	leveStr := ""
	if level == 1 {
		leveStr = ""
	} else {
		for i := level; i > 1; i-- {
			leveStr += "|-"
		}
	}

	files = append(files, leveStr+path)
	read, _ := ioutil.ReadDir(path)
	for _, f := range read {
		if f.IsDir() {
			newPath := path + "/" + f.Name()
			newLevel := level + 1
			files = addFilesX(newPath, files, newLevel)
		} else {
			newPath := path + "/" + f.Name()
			files = append(files, leveStr+newPath)
		}
	}
	return files
}
func main() {
	path := "/home/scorpio/app/golang-SDK/GOPATH/src/github.com"
	files := []string{}
	files = addFilesX(path, files, 1)
	for _, v := range files {
		fmt.Println(v)
	}
}
