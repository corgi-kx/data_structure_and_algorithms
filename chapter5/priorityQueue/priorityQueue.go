package priorityQueue

type store struct {
	serialNum int
	data interface{}
}

type priorityQueue struct {
	 data []store
	 length int
}

func New() *priorityQueue {
	p:=new(priorityQueue)
	return p
}

func (p *priorityQueue) Enqueue(priorityNum int,data interface{}) {
	newP:=store{priorityNum,data}
	p.data = append(p.data, newP)
	p.length ++
}

func (p *priorityQueue) heap() {
	mid:=p.length / 2
	for mid >=0 {
		leftchild,rightchild:=2*mid + 1 ,2*mid + 2
		if leftchild < p.length && p.data[leftchild].serialNum < p.data[mid].serialNum  {
			p.data[mid],p.data[leftchild] = p.data[leftchild],p.data[mid]

			//p.data[mid].data,p.data[leftchild].data = p.data[leftchild].data,p.data[mid].data
			//p.data[mid].serialNum,p.data[leftchild].serialNum = p.data[leftchild].serialNum,p.data[mid].serialNum
		}
		if rightchild < p.length && p.data[rightchild].serialNum < p.data[mid].serialNum {
			p.data[mid],p.data[rightchild] = p.data[rightchild],p.data[mid]
			//p.data[mid].data,p.data[rightchild].data = p.data[rightchild].data,p.data[mid].data
			//p.data[mid].serialNum,p.data[rightchild].serialNum = p.data[rightchild].serialNum,p.data[mid].serialNum
		}
		mid --
	}
}

func (p *priorityQueue) Dequeue() interface{} {
	if p.IsEmpty() {
		return  nil
	}else {
		p.heap()
		result := p.data[0]
		p.data = p.data[1:]
		p.length --
		return result
	}
}

func (p *priorityQueue) IsEmpty() bool {
	return p.length == 0
}
