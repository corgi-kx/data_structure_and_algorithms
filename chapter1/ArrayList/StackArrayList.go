package ArrayList

import "log"

type StackList struct {
	list *ArrayList
}

func NewStackList() *StackList {
	s := new(StackList)
	s.list = NewArrayList()
	return s
}

func (s *StackList) Clear_SK() {
	s = NewStackList()
}
func (s *StackList) Size_SK() int {
	return len(s.list.datastore)
}
func (s *StackList) Pop_SK() interface{} {
	if !s.IsEmpty_SK() {
		result := s.list.datastore[s.Size_SK()-1]
		s.list.datastore = s.list.datastore[:s.Size_SK()-1]
		return result
	} else {
		log.Panic("栈是空的")
		return nil
	}
}
func (s *StackList) Push_SK(i interface{}) {
	s.list.Append(i)
}
func (s *StackList) IsEmpty_SK() bool {
	return s.Size_SK() == 0
}

func (s *StackList) NewIterator() Iteratorer_SK {
	i := new(Iterator_SK)
	i.list = s
	return i
}
