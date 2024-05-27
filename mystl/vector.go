package mystl

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
)

type vectorer[T MyOrdered] interface {
	Sort(Cmp ...Comparator[T])          //利用比较器对其进行排序
	Size() (num int)                    //返回vector的长度
	Cap() (num int)                     //返回vector的容量
	Clear()                             //清空vector
	Empty() (b bool)                    //返回vector是否为空,为空则返回true反之返回false
	PushBack(e T) (err error)           //向vector末尾插入一个元素
	PopBack() (err error)               //弹出vector末尾元素
	Insert(idx uint64, e T) (err error) //向vector第idx的位置插入元素e,同时idx后的其他元素向后退一位
	Erase(idx uint64) (err error)       //删除vector的第idx个元素
	Reverse()                           //逆转vector中的数据顺序
	At(idx uint64) (e T)                //返回vector的第idx的元素
	Front() (e T)                       //返回vector的第一个元素
	Back() (e T)                        //返回vector的最后一个元素
	//Range() (start, end int)            //实现range遍历
	//Next() (e any, ok bool)    //返回下一个元素和是否还有下一元素
	Data() (m *[]T) //返回切片用于遍历
}

// Implementation

type Vector[T MyOrdered] struct {
	data []T
}

func NewVector[T MyOrdered](cap int, initial T) (v *Vector[T]) {
	v = &Vector[T]{
		data: make([]T, cap, cap),
	}
	if reflect.ValueOf(initial).IsZero() { //可以用反射, 不能用 initial.(type)
		return v
	}
	for i := 0; i < cap; i++ {
		v.data[i] = initial
	}
	return v
}

func NewVectorData[T MyOrdered](initData []T) (v *Vector[T]) {
	return &Vector[T]{
		data: initData,
	}
}

func NewVectorDefault[T MyOrdered]() (v *Vector[T]) {
	return &Vector[T]{
		data: make([]T, 0, 1),
	}
}

func (v *Vector[T]) Sort(Cmp ...Comparator[T]) {
	if v == nil {
		return
	}
	if len(Cmp) == 0 {
		//&& reflect.TypeOf(v.data).Implements(reflect.TypeOf((*cmp.Ordered)(nil)).Elem()) {
		// cannot use type cmp.Ordered outside a type constraint: interface contains type constraints
		slices.Sort(v.data)
	} else {
		slices.SortFunc(v.data, Cmp[0])
	}
}

func (v *Vector[T]) Size() (num int) {
	if v == nil {
		return 0
	}
	return len(v.data)
}

func (v *Vector[T]) Cap() (num int) {
	if v == nil {
		return 0
	}
	return cap(v.data)
}

func (v *Vector[T]) Clear() {
	if v == nil {
		fmt.Printf("Try to clear nil\n")
		return
	} else {
		v.data = make([]T, 0, 1)
	}
}

func (v *Vector[T]) Empty() (b bool) {
	if v == nil {
		return true
	}
	return v.Size() <= 0
}

func (v *Vector[T]) PushBack(e T) (err error) {
	if v == nil {
		return errors.New("error in PushBack, v is nil")
	}

	v.data = append(v.data, e)
	return nil
}

func (v *Vector[T]) PopBack() (err error) {
	if v == nil {
		return errors.New("error in PopBack, v is nil")
	}
	if v.Empty() {
		return errors.New("error in PopBack, v is empty")
	}
	v.data = v.data[:len(v.data)-1]
	return nil
}

func (v *Vector[T]) Insert(idx int, e T) (err error) {
	if v == nil {
		return errors.New("error in Insert, v is nil")
	}
	if idx < 0 || idx >= v.Size() {
		return fmt.Errorf("error in Insert, idx beyond range, idx:%v,v.Size:%d", idx, v.Size())
	}

	if v.Size() >= v.Cap() {
		v.PushBack(e)
	}
	var p int
	for p = v.Size() - 1; p > 0 && p > idx; p-- {
		v.data[p] = v.data[p-1]
	}
	v.data[p] = e
	return
}

func (v *Vector[T]) Erase(idx int) (err error) {
	if v == nil {
		return errors.New("error in Erase, v is nil")
	}
	if v.Empty() {
		return errors.New("error in Erase, v is empty")
	}
	if idx < 0 || idx >= v.Size() {
		return fmt.Errorf("error in Erase, vector idx(%v) <0 or >= size(%v)", idx, v.Size())
	}
	for p := idx; p < v.Size()-1; p++ {
		v.data[p] = v.data[p+1]
	}
	v.data = v.data[:v.Size()-1]
	return nil
}

func (v *Vector[T]) Reverse() {
	if v == nil {
		return
	}
	for i := 0; i < v.Size()/2; i++ {
		v.data[i], v.data[v.Size()-i-1] = v.data[v.Size()-i-1], v.data[i]
	}
}

func (v *Vector[T]) At(idx int) (e T) {
	if v == nil {
		//return e, errors.New("v is nil")
		panic("v is nil")
	}
	if idx < 0 || idx > v.Size() {
		panic(fmt.Sprintf("invalid idx:%v, size:%v", idx, v.Size()))
		//return e, errors.New(fmt.Sprintf("invalid idx:%q", idx))
	} //cannot return nil as T
	return v.data[idx]
}

func (v *Vector[T]) Front() (e T) {
	if v == nil {
		panic("v is nil")
	}
	if v.Size() == 0 {
		panic("v size is 0")
	}
	return v.data[0]
}

func (v *Vector[T]) Back() (e T) {
	if v == nil {
		panic("v is nil")
	}
	if v.Size() == 0 {
		panic("v size is 0")
	}
	return v.data[v.Size()-1]
}

func (v *Vector[T]) Data() *[]T {
	return &v.data
}
