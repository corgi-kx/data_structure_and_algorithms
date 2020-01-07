package CircleQueue

import "github.com/pkg/errors"

const queueCount = 5

type circleQueue struct {
	data  [queueCount]interface{}
	front int
	rear  int
}

func NewCircleQueue() *circleQueue {
	c := new(circleQueue)
	c.front = 0
	c.rear = 0
	c.data = [queueCount]interface{}{}
	return c
}

func (c *circleQueue) EnQueue(val interface{}) error {
	if (c.rear+1) % queueCount  == c.front % queueCount {
		return errors.New("循环队列已满,不能再添加了")
	}
	c.data[c.rear] = val
	c.rear = (c.rear + 1) % queueCount
	return nil
}

func (c *circleQueue) DeQueue() (interface{}, error) {
	if c.rear == c.front {
		return nil,errors.New("循环队列是空的！")
	}
	data:=c.data[c.front]
	c.data[c.front] = ""
	c.front = (c.front + 1)%queueCount
	return data,nil
}

func (c *circleQueue) GetLength() int {
	return (c.rear - c.front + queueCount) % queueCount
}

func (c *circleQueue) Clear() {
	c = NewCircleQueue()
}
