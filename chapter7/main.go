package main

import (
	"bufio"
	"data_structure_and_algorithms/chapter7/mergeMiddleWare"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	count        = 10000
	gap          = 10
	path  string = "D:/programming/golang/GOPATH/src/data_structure_and_algorithms/chapter7/"
)

func creatMIddleWare(fileName string, count, gap int) <-chan int {
	file, err := os.Create(path + fileName)
	if err != nil {
		panic(err)
	}
	sumSize := count * 8
	chunksize := sumSize / gap
	//生成随机数，写入文件
	mergeMiddleWare.WriteToFileCh(file, mergeMiddleWare.RandomArrs(sumSize))
	file.Close()

	//将文件进行读取并排序

	//将原需要排序的数据，进行分段排序
	fmt.Println("开始进行分段排序..........")
	mergeMiddleWare.SegmentedSort(path+fileName, path+"tmp.txt", sumSize, gap)
	fmt.Println("分段排序完成..........")

	//进行归并排序
	chs := []<-chan int{}
	fmt.Println("开始进行归并排序........")

	for i := 0; i < gap; i++ {
		tmpFile2, err := os.Open(path + "tmp.txt")
		if err != nil {
			panic(err)
		}
		tmpFile2.Seek(int64(i*chunksize), 0)
		ch := mergeMiddleWare.ReaderFromFile(bufio.NewReader(tmpFile2), chunksize)
		chs = append(chs, ch)
		//defer tmpFile2.Close()
	}
	return mergeMiddleWare.MergeN(chs...)
}

func showData(fileName string) <-chan int {
	file, err := os.Open(path + fileName)
	if err != nil {
		panic(err)
	}
	out := make(chan int)
	r := bufio.NewReader(file)
	go func() {
		for {
			buf := make([]byte, 8)
			_, err := r.Read(buf)
			if err != nil && err.Error() == "EOF" {
				break
			}
			out <- int(binary.BigEndian.Uint64(buf))
		}
		close(out)
	}()
	return out
}

func main() {
	//调用中间件
	ch := creatMIddleWare("old.txt", count, gap)
	//写入到新的文件
	file, _ := os.Create(path + "new.txt")
	mergeMiddleWare.WriteToFileCh(file, ch)
	//进行数据展示
	sch := showData("new.txt")
	for v := range sch {
		fmt.Println(v)
	}
}
