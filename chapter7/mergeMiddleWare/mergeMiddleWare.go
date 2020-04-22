package mergeMiddleWare

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"os"
	"sort"
)

//归并排序
func Merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func(ch1, ch2 <-chan int) {
		in1, v1 := <-ch1
		in2, v2 := <-ch2
		for v1 || v2 {
			if !v2 || (v1 && in1 <= in2) {
				out <- in1
				in1, v1 = <-ch1
			} else {
				out <- in2
				in2, v2 = <-ch2
			}
		}
		close(out)
	}(ch1, ch2)
	return out
}

func MergeN(chs ...<-chan int) <-chan int {
	if len(chs) == 1 {
		return chs[0]
	} else {
		mid := len(chs) / 2
		return Merge(MergeN(chs[:mid]...), MergeN(chs[mid:]...))
	}
}

func RandomArrs(nums int) <-chan int {
	ch := make(chan int, 1024)
	go func() {
		for i := 0; i < nums; i++ {
			nums, err := rand.Int(rand.Reader, big.NewInt(256))
			if err != nil {
				panic(err)
			}
			ch <- int(nums.Int64())
		}
		close(ch)
	}()

	return ch
}

func SegmentedSort(sourcePath, tmpPath string, sumSize, gap int) {
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	//将文件进行读取并排序
	chunksize := sumSize / gap
	for i := 0; i < gap; i++ {
		file2, _ := os.Open(sourcePath)
		file2.Seek(int64(i*chunksize), 0)
		defer file2.Close()
		ch := ReaderFromFile(file2, chunksize)
		wch := IntsSortsToChan(ch)
		for v := range wch {
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, uint64(v))
			tmpFile.Write(buf)
		}

	}
}

func IntsSortsToChan(ch <-chan int) <-chan int {
	out := make(chan int, 1024)
	nums := []int{}
	for v := range ch {
		nums = append(nums, v)
	}
	sort.Ints(nums)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func WriteToFileCh(writer io.Writer, ch <-chan int) {
	w := bufio.NewWriter(writer)
	for v := range ch {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		w.Write(buf)
	}
	w.Flush()
}

func ReaderFromFile(infile io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	r := bufio.NewReader(infile)
	go func() {
		result := 0
		for {
			data := make([]byte, 8)
			n, err := r.Read(data)
			if err != nil && err.Error() == "EOF" {
				break
			}
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(data))
			}
			result += n
			if result >= chunkSize {
				break
			}
		}
		close(out)
	}()

	return out
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}
