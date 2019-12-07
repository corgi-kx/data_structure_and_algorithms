package ArrayList

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type ArrayList struct {
	datastore []interface{}
	size int
}

func NewArrayList() *ArrayList {
	return &ArrayList{make([]interface{},0,10),0}
}

func (a *ArrayList) Get(index int) (interface{},error) {
	err:=a.checkIndexOut(index)
	if err != nil {
		return nil,err
	}
	return a.datastore[index],nil
}

func (a *ArrayList) Set(index int,newval interface{}) error {
	err:=a.checkIndexOut(index)
	if err != nil {
		return err
	}
	a.datastore[index] = newval
	return nil
}

func (a *ArrayList) Delete(index int) error {
	err:=a.checkIndexOut(index)
	if err != nil {
		return err
	}
	a.datastore = append(a.datastore[:index],a.datastore[index + 1 :]...)
	a.size --
	return nil
}

func (a *ArrayList)Clear() {
	a.datastore = make([]interface{},0,10)
	a.size = 0
}

func (a *ArrayList)	Append(newval interface{}) {
	a.datastore = append(a.datastore,newval)
	a.size ++
}

func (a *ArrayList) Insert(index int,newval interface{}) error {
	err:=a.checkIndexOut(index)
	if err != nil {
		return err
	}
	a.datastore = append(a.datastore, newval)

	for i:=len(a.datastore) - 1;i>=index;i-- {
		a.datastore[i] = a.datastore[i - 1]
	}
	a.datastore[index] = newval
	a.size ++
	return nil
}

func (a *ArrayList) Iterator() Iteratorer {
	it:=new(Iterator)
	it.list = a
	it.currentIndex = 0
	return it
}


func (a *ArrayList) String() string {
		return fmt.Sprint(a.datastore)
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) checkIndexOut(index int) error {
	if index < 0 || index > a.size {
		return errors.New("数组下标越界,当前数组长度最大值为"+strconv.Itoa(index))
	}else {
		return nil
	}
}