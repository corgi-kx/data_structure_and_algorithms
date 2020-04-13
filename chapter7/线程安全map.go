package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SysMap struct{
	m map[int]string
	mutex *sync.RWMutex
}

func (s *SysMap) write(num int,str string) {
	s.mutex.Lock()
	defer  s.mutex.Unlock()
	s.m[num] = str
}

func (s *SysMap) Read()  {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for k,v:=range s.m {
		fmt.Println(k,v)
	}
}

func New() *SysMap {
	s:=SysMap{make(map[int]string),new(sync.RWMutex)}
	return &s
}
func main() {
	s:=New()
	wait:= sync.WaitGroup{}
	wait.Add(10)
	for i:=0;i<10;i++ {
		go func(num int,sysMap *SysMap) {
			sysMap.write(num,strconv.Itoa(num))
			wait.Done()
		}(i,s)
	}
	wait.Wait()
	s.Read()
}
