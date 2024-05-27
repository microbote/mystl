package mystl

type listNode[T Load] struct {
	next *listNode[T]
	prev *listNode[T]
	data T
}

func newListNode[T Load](e T) (n *listNode[T]) {
	return &listNode[T]{
		data: e,
		prev: nil,
		next: nil,
	}
}

//=========================== Lister ====================================

type Lister[T Load] interface {
	Begin() (i Iterator[T])
	End() (i Iterator[T])
	Tail() (i Iterator[T])
	Iterator() (i Iterator[T])  //创建一个包含链表中所有元素的迭代器并返回其指针
	Sort(Cmp ComparatorAny[T])  //将链表中所承载的所有元素进行排序
	Size() (size uint64)        //返回链表所承载的元素个数
	Clear()                     //清空该链表
	Empty() (b bool)            //判断该链表是否位空
	Insert(it Iterator[T], e T) //向链表的idx位(下标从0开始)插入元素组e
	Erase(it Iterator[T])       //删除第idx位的元素(下标从0开始)
	Get(it Iterator[T]) (e T)   //获得下标为idx的元素
	Set(it Iterator[T], e T)    //在下标为idx的位置上放置元素e
	PushBack(e T)
	PopBack()
	Back() (e T)
	Front() (e T)
	PushFront(e T)
	PopFront()
}

// =========================== ListIterator ====================================
type ListIterator[T Load] struct {
	list Lister[T] //to the list
	cur  *listNode[T]
}

func (it *ListIterator[T]) Begin() (i Iterator[T]) { //将该迭代器设为位于首节点并返回新迭代器
	if it == nil || it.list == nil {
		return nil
	}
	return it.list.Begin()
}
func (it *ListIterator[T]) Last() (i Iterator[T]) { //将该迭代器设为位于尾节点并返回新迭代器
	if it == nil || it.list == nil {
		return nil
	}
	return it.list.Tail()
}

func (it *ListIterator[T]) End() (i Iterator[T]) {
	if it == nil || it.list == nil {
		return nil
	}
	return it.list.End()
}
func (it *ListIterator[T]) Value() (e T) { //返回该迭代器下标所指元素
	if it == nil {
		panic("null ptr to list iterator")
	}
	return it.cur.data
}
func (it *ListIterator[T]) CurRef() (e *T) { //
	if it == nil {
		panic("null ptr to list iterator")
	}
	return &it.cur.data
}
func (it *ListIterator[T]) HasNext() (b bool) { //判断该迭代器是否可以后移
	if it == nil {
		panic("null ptr to list iterator")
	}
	return it.NotEqual(it.list.End())
}
func (it *ListIterator[T]) Next() (b bool) { //将该迭代器后移一位
	if it == nil {
		panic("null ptr to list iterator")
	}
	if !it.HasNext() {
		return false
	}
	it.cur = it.cur.next
	return true
}
func (it *ListIterator[T]) HasPre() (b bool) { //判断该迭代器是否可以前移
	if it == nil {
		panic("null ptr to list iterator")
	}
	if it.Equal(it.list.Begin()) {
		return false
	}
	return true
}
func (it *ListIterator[T]) Pre() (b bool) { //将该迭代器前移一位
	if it == nil {
		panic("null ptr to list iterator")
	}
	if !it.HasPre() {
		return false
	}
	it.cur = it.cur.prev
	return true
}
func (it *ListIterator[T]) Equal(i Iterator[T]) (b bool) { //是否相同
	if it == nil {
		panic("null ptr to list iterator")
	}
	oth := i.(*ListIterator[T])
	if it.list == oth.list && it.cur == oth.cur {
		return true
	} else {
		return false
	}
}

func (it *ListIterator[T]) NotEqual(i Iterator[T]) (b bool) { //是否不同
	if it == nil {
		panic("null ptr to list iterator")
	}
	return !it.Equal(i)
}

// =========================== List ====================================
type List[T Load] struct {
	node *listNode[T]
	size int
}

func NewList[T Load]() (l *List[T]) {
	var e T
	l = &List[T]{
		node: newListNode(e),
		size: 0,
	}
	l.node.prev = l.node
	l.node.next = l.node
	return
}


