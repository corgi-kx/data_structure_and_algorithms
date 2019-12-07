package Stack

import (
	"log"
)

type Stack interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(i interface{})
	IsEmpty() bool
}

type StackList struct {
	datastore []interface{}
}

func NewStackList() *StackList {
	s:= new(StackList)
	s.datastore = make([]interface{},0)
	return s
}

func (s *StackList) Clear() {
	s = NewStackList()
}
func (s *StackList) Size() int {
	return len(s.datastore)
}
func (s *StackList) Pop() interface{} {
	if !s.IsEmpty() {
		result := s.datastore[s.Size() - 1]
		s.datastore = s.datastore[:s.Size() - 1]
		return result
	}else {
		log.Panic("栈是空的")
		return nil
	}
}
func (s *StackList) Push(i interface{})  {
	s.datastore = append(s.datastore, i)
}
func (s *StackList) IsEmpty() bool {
	return s.Size() == 0
}