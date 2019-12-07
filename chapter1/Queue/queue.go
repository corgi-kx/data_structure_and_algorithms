package Queue


type Queuer interface {
	DeQueue(interface{})
	EnQueue()interface{}
	Size() uint
	Front()interface{}
	Last()interface{}
	IsEmpty() bool
	Clear()
}

type Queue struct {
	datastore []interface{}
	size uint
}

func NewQueue() *Queue {
	q:=new(Queue)
	q.datastore = make([]interface{},0)
	q.size = 0
	return q
}

func (q *Queue) EnQueue(val interface{}) {
	q.datastore = append(q.datastore, val)
	q.size ++
}
func (q *Queue) DeQueue() interface{}{

	if q.IsEmpty() {
		return nil
	}
		data:=q.datastore[0]
		q.datastore = q.datastore[1:q.size]
		q.size --
		return data
}
func (q *Queue) Size() uint{
	return q.size
}
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.datastore[0]
}

func (q *Queue) Last() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.datastore[q.size - 1]
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}else {
		return false
	}
}
func (q *Queue) Clear() {
	q = NewQueue()
}