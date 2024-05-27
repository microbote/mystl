package mystl

type BinarySearchTreer[T Load] interface {
	Iterator() (i *Iterator[T])
	Size() (num int)                 //返回该二叉树中保存的元素个数
	Clear()                          //清空该二叉树
	Empty() (b bool)                 //判断该二叉树是否为空
	Insert(e T) (b bool)             //向二叉树中插入元素e
	Erase(e T) (b bool)              //从二叉树中删除元素e
	Count(e T) (num int)             //从二叉树中寻找元素e并返回其个数
	Find(e T) (i *Iterator[T])       //查找
	LowerBound(e T) (i *Iterator[T]) //>=e
	UpperBound(e T) (i *Iterator[T]) //>e
}
