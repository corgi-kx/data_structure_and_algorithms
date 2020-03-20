package main

import "fmt"

type  Element int64
type LinkNode struct {
	data Element
	pNext *LinkNode
}


type Linker interface {
	Add(data Element)
	HeadInsert(data Element)
	Append(data Element)
	Show()
	IsEmpty() bool
	GetLength() int
	Search(data Element) bool
	GetData(num int) Element
	Delete(data Element) bool
	Clear()
	Reversal()
	Insert(id int,data Element)
	BubbleSort()
	SelectSort()
	InsertSort()
	MergeSort()
	QuickSort()
}

func New() *LinkNode {
	return &LinkNode{0,nil}
}
func (l *LinkNode) HeadInsert(data Element) {
	nNode := New()
	nNode.data = data
	if l.IsEmpty() {
		l.pNext = nNode
	}else  {
		nNode.pNext = l.pNext
		l.pNext = nNode
	}
}

func (l *LinkNode) Append(node *LinkNode) {
	if l.IsEmpty() {
		l.pNext = node
	}else {
		lastNode := l.pNext
		for lastNode.pNext != nil {
			lastNode = lastNode.pNext
		}
		lastNode.pNext = node
	}
}

func (l *LinkNode) Add(data Element) {
	nNode := New()
	nNode.data = data
	l.Append(nNode)
}

func (l *LinkNode) Show() {
	if l.IsEmpty() {
		fmt.Println("link is nil")
		return
	}
	cNode:=l.pNext
	fmt.Println(cNode)
	for cNode.pNext !=nil {
		cNode = cNode.pNext
		fmt.Println(cNode)
	}
	fmt.Println("show over")
}
func (l *LinkNode) IsEmpty() bool{
	return l.pNext == nil
}


func (l *LinkNode) GetLength() int{
	if l.IsEmpty() {
		return 0
	}else {
		lastNode := l.pNext
		length:=1
		for lastNode.pNext != nil {
			lastNode = lastNode.pNext
			length ++
		}
		return length
	}
}

func (l *LinkNode) Search(data Element) bool {
	if l.IsEmpty() {
		return false
	}else {
		lastNode := l.pNext
		if lastNode.data == data  {
			return true
		}
		for lastNode.pNext != nil {
			lastNode = lastNode.pNext
			if lastNode.data == data  {
				return true
			}
		}
		return false
	}
}

func (l *LinkNode) GetData(num int) Element{
	if num < 0 || l.IsEmpty() {
		panic("链表为空或者传入的数字太小")
	}else {
		lastNode := l
		id := 0
		for lastNode.pNext != nil {
			lastNode = lastNode.pNext
			id ++
			if id == num {
				return lastNode.data
			}

		}
	}
	panic("没有找到！！！")
}

func (l *LinkNode) Delete(data Element) bool {
	if l.IsEmpty() {
		return false
	}else {
		lastNode := l
		for lastNode.pNext != nil {
			if lastNode.pNext.data == data {
				lastNode.pNext = lastNode.pNext.pNext
				return true
			}
			lastNode = lastNode.pNext
		}
	}
	return false
}

func (l *LinkNode) Clear() {
	l.pNext = nil
}


func (l *LinkNode) Insert(id int,data Element) {
	if l.IsEmpty() || l.GetLength() < id{
		return
	}else {
		nNode:=New()
		nNode.data = data

		lastNode := l
		num:=0
		for lastNode.pNext != nil {
			lastNode = lastNode.pNext
			num ++
			if num == id {
				nNode.pNext = lastNode.pNext
				lastNode.pNext = nNode
			}
		}
	}
}

func (l *LinkNode) Reversal() {
	if l.IsEmpty() {
		return
	}
	lastnode := l.pNext
	for lastnode != nil {
		cNode:=lastnode.pNext
		if cNode == nil {
			return
		}
		lastnode.pNext = lastnode.pNext.pNext
		cNode.pNext = l.pNext
		l.pNext = cNode
	}
}

func (l *LinkNode) BubbleSort() {
	if l.IsEmpty() {
		return
	}
	for iNode:=l.pNext;iNode!=nil;iNode = iNode.pNext {
		for jNode:=l.pNext;jNode!=nil;jNode = jNode.pNext {
			if jNode.pNext !=nil && jNode.data > jNode.pNext.data {
				jNode.data,jNode.pNext.data = jNode.pNext.data,jNode.data
			}
		}
	}
}

func (l *LinkNode) SelectSort(){
	if l.IsEmpty() {
		return
	}
	for iNode:=l.pNext;iNode != nil;iNode = iNode.pNext {
		tmpNode :=iNode
		for jNode:=iNode;jNode!=nil;jNode = jNode.pNext {
			if jNode.pNext !=nil && tmpNode.data > jNode.pNext.data {
				tmpNode = jNode.pNext
			}
		}
		if tmpNode != iNode {
			iNode.data,tmpNode.data = tmpNode.data,iNode.data
		}
	}
}

func (l *LinkNode) InsertSort() {
	if l.IsEmpty() {
		return
	}
	pStart,pEnd:=l,l.pNext
	for iNode:=pEnd.pNext;iNode!=nil;iNode = pEnd.pNext {
		pStart = l
		for pStart.pNext != nil && pStart != pEnd.pNext && iNode.data > pStart.pNext.data{
			pStart = pStart.pNext
		}
		pEnd.pNext = iNode.pNext
		iNode.pNext = pStart.pNext
		pStart.pNext = iNode
		if pEnd.pNext == nil {
			break
		}

	}
}

func (l *LinkNode) MergeSort() {
	l.pNext = MergeSortIn(l)
}

func MergeSortIn(node *LinkNode) *LinkNode{
	if node == nil || node.pNext == nil || node.pNext.pNext == nil {
		return node
	}
	mid,end:=node.pNext,node.pNext
	pre:=New()
	for end.pNext != nil && end.pNext.pNext != nil {
		end = end.pNext.pNext
		pre = mid
		mid = mid.pNext
	}
	if node.GetLength() == 2 {
		mid = node.pNext.pNext
		node.pNext.pNext = nil
	}
	pre.pNext = nil
	finalMid:=New()
	finalMid.pNext = mid
	//node.Show()
	//finalMid.Show()
	//return nil
	left:= MergeSortIn(node)
	right:= MergeSortIn(finalMid)
	return Merge(left,right)
}

func Merge(left,right *LinkNode) *LinkNode{
	result:=New()
	for left != nil && right != nil {
		if left.data < right.data {
			result.Add(left.data)
			left = left.pNext
		}else {
			result.Add(right.data)
			right = right.pNext
		}
	}
	for  left != nil {
		result.Add(left.data)
		left = left.pNext
	}
	for right != nil {
		result.Add(right.data)
		right = right.pNext
	}
	return result
}



func (l *LinkNode) QuickSort() {
		l.pNext = quickSortFuc(l.pNext)
}

func quickSortFuc(node *LinkNode) *LinkNode {
	if node == nil || node.pNext == nil{
		return node
	}
	mid := node

	left,right := New(),New()
	for node != nil {
		if node.data <= mid.data && node != mid {
			left.Add(node.data)
		}else if node.data > mid.data && node != mid{
			right.Add(node.data)
		}
		node = node.pNext
	}

	mid.pNext = nil

	leftResult,rightResult := New(),New()
	if left.pNext != nil {
		leftResult.pNext =quickSortFuc(left.pNext)
	}
	if right.pNext!= nil {
		rightResult.pNext=quickSortFuc(right.pNext)
	}


	final :=New()
	if leftResult.pNext!=nil {
		for leftResult.pNext!= nil {
			final.Add(leftResult.pNext.data)
			leftResult = leftResult.pNext
		}
	}
	final.Add(mid.data)
	if rightResult.pNext != nil{
		for rightResult.pNext!= nil {
			final.Add(rightResult.pNext.data)
			rightResult = rightResult.pNext
		}
	}

	return final.pNext
}


func main() {
	link:=New()
	link.Add(9)
	link.Add(3)
	link.Add(0)
	link.Add(1)
	link.Add(8)
	link.Add(4)
	link.Add(5)
	link.Add(2)
	//link.HeadInsert(2)
	//link.HeadInsert(1)
	//
	//link.Show()
	//
	//link.Delete(3)
	//link.Insert(2,3)
	//link.Show()
	//link.Reversal()
	link.QuickSort()
	link.Show()
	//link.Reversal()
	//link.Show()
}