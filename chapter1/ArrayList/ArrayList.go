package ArrayList

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type List interface {
	Get(index int) (interface{},error)
	Set(index int,newval interface{}) error
	Delete(index int) error
	Clear()
	Append(newval interface{})
	Insert(index int,newval interface{}) error
}

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
	a.datastore = append(a.datastore[:index],a.datastore[index + 1 :])
	a.size --
	return nil
}

func (a *ArrayList)Clear() {
	a.datastore = make([]interface{},10)
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
	a.datastore = append(a.datastore, "")

	for i:=len(a.datastore);i>=index;i-- {
		fmt.Println(a.size)
		a.datastore[i-1] = a.datastore[i -2]
	}
	a.datastore[index] = newval
	a.size ++
	return nil
}



func (a *ArrayList) checkIndexOut(index int) error {
	if index < 0 || index > a.size {
		return errors.New("数组下标越界,当前数组长度最大值为"+strconv.Itoa(index))
	}else {
		return nil
	}
}