package mystl

type rbtreer[T load] interface {
	Iterator() (i *Iterator.Iterator) //返回包含该二叉树的所有元素,重复则返回多个
	Size() (num int)                  //返回该二叉树中保存的元素个数
	Clear()                           //清空该二叉树
	Empty() (b bool)                  //判断该二叉树是否为空
	Insert(e interface{}) (b bool)    //向二叉树中插入元素e
	Erase(e interface{}) (b bool)     //从二叉树中删除元素e
	Count(e interface{}) (num int)    //从二叉树中寻找元素e并返回其个数
}
