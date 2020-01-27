package link

import (
	"fmt"
	"log"
)

type LinkLister interface {
	Add(data interface{})
	InsertFront(data interface{})
	InsertRear(data interface{})
	InsertIndex(index int, data interface{})
	ReplayNode(target, newNode *LinkNode)
	UpdateByIndex(index int, data interface{})
	DeleteByNode(node *LinkNode)
	DeleteByIndex(index int)
	GetFirstNode() *LinkNode
	GetRearNode() *LinkNode
	GetMidNode() *LinkNode
	GetSize() int
	IsEmpt() bool
	Reversal()
	String() string
}

type LinkList struct {
	head *LinkNode
	size int
}

func NewLinkList() *LinkList {
	l := new(LinkList)
	l.head = NewLinkNode(nil)
	l.size = 0
	return l
}
func (l *LinkList) InsertIndex(index int, data interface{}) {
	newNode := NewLinkNode(data)
	node := l.head
	nodeIndex := 0
	for node != nil {
		if nodeIndex == index {
			newNode.next = node.next
			node.next = newNode
			l.size++
			return
		}
		node = node.next
		nodeIndex++
	}
	log.Panic("InsertIndex err: 下标在当前链表中不存在")
}

func (l *LinkList) InsertFront(data interface{}) {
	node := NewLinkNode(data)
	if l.IsEmpt() {
		l.head.next = node
		l.size++
	} else {
		node.next = l.head.next
		l.head.next = node
		l.size++
	}
}

func (l *LinkList) InsertRear(data interface{}) {
	l.Add(data)
}

func (l *LinkList) Add(data interface{}) {
	if l.IsEmpt() {
		node := NewLinkNode(data)
		l.head.next = node
		l.size++
	} else {
		node := NewLinkNode(data)
		lastNode := l.GetRearNode()
		lastNode.next = node
		l.size++
	}
}
func (l *LinkList) GetRearNode() *LinkNode {
	if !l.IsEmpt() {
		rear := l.head.next
		for rear.next != nil {
			rear = rear.next
		}
		return rear
	}
	return nil
}

func (l *LinkList) GetFirstNode() *LinkNode {
	if !l.IsEmpt() {
		return l.head.next
	} else {
		return nil
	}
}

func (l *LinkList) GetMidNode() *LinkNode {
	if !l.IsEmpt() {
		one := l.head.next
		two := l.head.next
		for two.next != nil && two.next.next != nil {
			two = two.next.next
			one = one.next
		}
		return one
	} else {
		return nil
	}
}

func (l *LinkList) ReplayNode(target, newNode *LinkNode) {
	if l.IsEmpt() {
		log.Panic("UpdateByNode err : 链表是空的")
	}
	node := l.head
	for node.next != nil {
		if target == node.next {
			newNode.next = node.next.next
			node.next = newNode
		}
		node = node.next
	}
}
func (l *LinkList) UpdateByIndex(index int, data interface{}) {
	if l.IsEmpt() || l.size-1 < index {
		log.Panic("UpdateByIndex err: 下标不存在")
	}
	nodeIndex := 0
	node := l.head.next
	for node != nil {
		if index == nodeIndex {
			node.data = data
			return
		}
		node = node.next
		nodeIndex++
	}
}

func (l *LinkList) DeleteByNode(targetNode *LinkNode) {
	if l.IsEmpt() {
		log.Panic("DeleteByNode err : 链表是空的")
	}
	node := l.head
	for node.next != nil {
		if targetNode == node.next {
			node.next = node.next.next
			l.size--
			return
		}
		node = node.next
	}
}
func (l *LinkList) DeleteByIndex(index int) {
	if l.IsEmpt() || l.size-1 < index {
		log.Panic("DeleteByNode err : 下标不存在")
	}
	nodeIndex := 0
	node := l.head
	for node.next != nil {
		if index == nodeIndex {
			node.next = node.next.next
			l.size--
			return
		}
		node = node.next
		nodeIndex++
	}
}

func (l *LinkList) Reversal() {
	if l.IsEmpt() {
		log.Panic("Reversal err : 链表是空的")
	}

	head,index := l.head.next,l.head.next
	for index.next!=nil {
		rear := index.next
		index.next = index.next.next
		head.next = index.next
		rear.next = l.head.next
		l.head.next = rear
	}

	//cur := head
	//var pre *ListNode = nil
	//for cur != nil {
	//	pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	//}
	//return pre
}

func (l *LinkList) IsEmpt() bool {
	if l.head.next == nil {
		return true
	} else {
		return false
	}
}
func (l *LinkList) GetSize() int {
	return l.size
}
func (l *LinkList) String() string {
	if l.head.next != nil {
		s := "["
		node := l.head.next
		for node != nil {
			if node.next != nil {
				s += fmt.Sprint(node.data) + ","
			} else {
				s += fmt.Sprint(node.data) + "]"
			}
			node = node.next
		}
		return s
	} else {
		return ""
	}
}
