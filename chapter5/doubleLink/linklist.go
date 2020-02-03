package doubleLink

import (
	"fmt"
	"log"
)

type DoubleLinker interface {
	String() string
	ReversePrint() string
	InsertFront(data interface{})
	InsertIndex(index int , data interface{})
	Add(data interface{})
	Update(index int,data interface{})
	Delete(index int)
	GetSize() int
	IsEmpty() bool

}

type DoubleLinkList struct {
	header *node
	size 	int
}

func NewDoubleLinkList() *DoubleLinkList{
	l:=new(DoubleLinkList)
	l.header = NewNode(nil)
	l.size = 0
	return  l
}
func  (l *DoubleLinkList) InsertFront(data interface{}) {
	newNode := NewNode(data)
	if l.IsEmpty() {
		l.header.nextNode = newNode
		newNode.preNode = l.header
	}else  {
		head := l.header.nextNode
		newNode.nextNode,head.preNode = head,newNode
		l.header.nextNode = newNode
	}
	l.size ++
}

func  (l *DoubleLinkList) InsertIndex(index int , data interface{}) {
	if index  >= l.size {
		log.Panic("InsertIndex err : 无此下标" )
	}
	newNode := NewNode(data)
	if index == 0 {
		node := l.header.nextNode
		newNode.nextNode = node
		node.preNode = newNode
		l.header.nextNode = newNode
	}else  {
		head := l.header.nextNode
		for index > 1 {
			head = head.nextNode
			index --
		}
		newNode.nextNode = head.nextNode
		head.nextNode.preNode = newNode

		head.nextNode = newNode
		newNode.preNode = head
	}
	l.size ++
}

func (l *DoubleLinkList) Add (data interface{}) {
	newNode := NewNode(data)
	if l.IsEmpty() {
		l.header.nextNode = newNode
		newNode.preNode = l.header
	}else  {
		node := l.header.nextNode
		for node.nextNode != nil {
			node = node.nextNode
		}
		node.nextNode = newNode
		newNode.preNode = node
	}
	l.size ++
}

func (l *DoubleLinkList) IsEmpty () bool {
	if l.header.nextNode == nil {
		return true
	}else  {
		return false
	}
}

func (l *DoubleLinkList) GetSize() int {
	return l.size
}

func (l *DoubleLinkList) String() string {
	if !l.IsEmpty() {
		s := "["
		node := l.header.nextNode
		for node != nil {
			if node.nextNode != nil {
				s += fmt.Sprint(node.data)  + ","
			} else {
				s += fmt.Sprint(node.data) + "]"
			}
			node = node.nextNode
		}
		return s
	} else {
		return ""
	}
}

func (l *DoubleLinkList) ReversePrint() string {
	if l.IsEmpty() {
		return ""
	}else {
		node := l.header.nextNode
		for node.nextNode != nil {
			node = node.nextNode
		}
		s := "["
		for node != nil {
			if node.preNode != nil {
				s += fmt.Sprint(node.data)  + ","
			} else {
				s += fmt.Sprint(node.data) + "]"
			}
			node = node.preNode
		}
		return s
	}
}