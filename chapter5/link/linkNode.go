package link

type LinkNode struct {
	data interface{}
	next *LinkNode
}

func NewLinkNode(data interface{}) *LinkNode{
	l := new(LinkNode)
	l.data = data
	return l
}

func (l *LinkNode) value () interface{}  {
	return l.data
}