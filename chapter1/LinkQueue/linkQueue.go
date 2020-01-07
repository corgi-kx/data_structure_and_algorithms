package LinkeQueue

import "github.com/pkg/errors"

type node struct {
	data interface{}
	next *node
}

type queue struct {
	front *node
	rear *node
	size int
}

func New() *queue {
	return new(queue)
}

func (q *queue) EnQueue(val interface{}) {
	newNode := &node{val,nil}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else  {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size ++
}


func (q *queue) DeQueue() (interface{},error){
	if q.front == nil {
		return nil,errors.New("链式队列是空的")
	}
	data := q.front.data
	q.front = q.front.next
	q.size --
	return data,nil
}

func (q *queue) Size() int{
	return q.size
}


func (q *queue) IsEmpty() bool{
	if q.front == nil {
		return true
	}else {
		return false
	}
}