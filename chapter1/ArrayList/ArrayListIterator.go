package ArrayList

import log "github.com/corgi-kx/logcustom"

type Iterator struct {
	list Lister
	currentIndex int
}

func (a *Iterator) HasNext() bool {
	return a.currentIndex  < a.list.Size()
}

func (a *Iterator)  Next() interface{} {
	val,err:= a.list.Get(a.currentIndex)
	if err != nil {
		log.Panic(err)
	}
	a.currentIndex ++
	return val
}

func (a *Iterator)  Remove()  {
	a.currentIndex --
	err:=a.list.Delete(a.currentIndex)
	if err != nil {
		log.Panic(err)
	}
}

func (a *Iterator)  GetIndex() int {
	return a.currentIndex
}