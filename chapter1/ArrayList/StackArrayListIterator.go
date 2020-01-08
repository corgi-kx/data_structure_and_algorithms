package ArrayList

type Iterator_SK struct {
	list Stack
}

func (a *Iterator_SK) HasNext_SK() bool {
	return a.list.Size_SK() > 0
}

func (a *Iterator_SK) Next_SK() interface{} {
	return a.list.Pop_SK()
}

func (a *Iterator_SK) GetIndex_SK() int {
	return a.list.Size_SK()
}
