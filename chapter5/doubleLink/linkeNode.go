package doubleLink


type node struct {
	data interface{}
	preNode *node
	nextNode *node
}

func NewNode(data interface{}) *node{
	n:= new(node)
	n.data = data
	n.nextNode = nil
	n.preNode = nil
	return n
}