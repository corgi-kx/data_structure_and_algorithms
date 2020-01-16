package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

var (
	idmax      int64 = 1000000000
	idmin      int64 = idmax / 10
	passmin    int64 = 8
	passmax    int64 = 20
	threads          = 10
	dataLength       = 30000000
)

var dictionary = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'j', 'h', 'i', 'g', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'J', 'H', 'I', 'G', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'!', '@', '#', '$', '%', '&', '*', '?', '=',
}

func main() {
	now := time.Now()
	f, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ch := make(chan string, dataLength/10)
	for i := 0; i < threads; i++ {
		go RandomGen(ch)
	}

	w := bufio.NewWriter(f)
	for i := 0; i < dataLength; i++ {
		_, err := w.WriteString(<-ch + "\n")
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
	fmt.Println("用时：", time.Since(now))
}

func RandomGen(ch chan string) {
	for {
		var id int64
		var password string
		//生成9位数id
		max := big.NewInt(idmax)
		for {
			randNum, _ := rand.Int(rand.Reader, max)
			if randNum.Int64() > idmin {
				id = randNum.Int64()
				break
			}
		}

		//生成随机密码
		finalLength := big.NewInt(0)
		for {
			finalLength, _ = rand.Int(rand.Reader, big.NewInt(passmax))
			if finalLength.Int64() >= passmin {
				break
			}
		}

		dl := len(dictionary)
		pass := []byte{}
		for i := 0; i < int(finalLength.Int64()); i++ {
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(dl)))
			pass = append(pass, dictionary[index.Int64()])
		}
		password = string(pass)

		str := strconv.Itoa(int(id)) + " # " + password
		ch <- str
	}

}
