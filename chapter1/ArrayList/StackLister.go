package ArrayList


type Stack interface {
	Clear_SK()
	Size_SK() int
	Pop_SK() interface{}
	Push_SK(i interface{})
	IsEmpty_SK() bool
	NewIterator() Iteratorer_SK
}



type Iteratorer_SK interface {
	//判断是否迭代完毕
	HasNext_SK() bool
	//根据下标获取当前数据
	Next_SK() interface{}
	//获取当前迭代下标
	GetIndex_SK() int
}