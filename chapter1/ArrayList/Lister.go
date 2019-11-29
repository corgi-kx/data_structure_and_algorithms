package ArrayList

type Lister interface {
	//根据索引获取下标所在的数组数据
	Get(index int) (interface{}, error)
	//根据索引替换下标所在的数组数据
	Set(index int, newval interface{}) error
	//根据索引删除下标所在的数据数据
	Delete(index int) error
	//清空数组（利用了gc特性）
	Clear()
	//追加数组数据
	Append(newval interface{})
	//在指定索引后面，插入新数据
	Insert(index int, newval interface{}) error
	//返回数组信息
	String() string
	//返回数组长度
	Size() int
	//获取迭代器
	Iterator() Iteratorer
}

type Iteratorer interface {
	//判断是否迭代完毕
	HasNext() bool
	//根据下标获取当前数据
	Next() interface{}
	//删除数组中当前迭代数据
	Remove()
	//获取当前迭代下标
	GetIndex() int
}
