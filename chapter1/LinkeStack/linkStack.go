package LinkeStack

import "github.com/pkg/errors"

type node struct {
	data interface{}
	next *node
	size int
}

func New() *node {
	return new(node)
}


func (n *node) Push(val interface{}) {
	newNode := New()
	newNode.data = val
	newNode.next = n.next
	n.next = newNode
	n.size ++
}

func (n *node) Pop() (interface{},error){
	if n.IsEmpty() {
		return nil,errors.New("栈是空的")
	}
	data := n.next.data
	n.next = n.next.next
	n.size --
	return data,nil
}

func (n *node) Size() int{
	return n.size
}

func (n *node) IsEmpty()bool {
	if n.next == nil {
		return true
	}else  {
		return false
	}
}