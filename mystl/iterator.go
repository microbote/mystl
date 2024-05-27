package mystl

//@Title		Iterator
//@Description
//		迭代器
//		定义了一套迭代器接口和迭代器类型
//		本套接口定义了迭代器所要执行的基本函数
//		数据结构在使用迭代器时需要重写函数
//		其中主要包括:生成迭代器,移动迭代器,判断是否可移动

// Iterator迭代器接口
// 定义了一套迭代器接口函数
// 函数含义详情见下列描述
type Iterator interface {
	Begin() (i Iterator) //将该迭代器设为位于首节点并返回新迭代器
	Last() (i Iterator)  //将该迭代器设为位于尾节点并返回新迭代器
	End() (i Iterator)   // 左闭右开

	Value() (e any) //返回该迭代器下标所指元素
	Ptr() (e *any)
	HasNext() (b bool) //判断该迭代器是否可以后移
	Next() (b bool)    //将该迭代器后移一位

	Equal(i Iterator) (b bool)    //是否相同
	NotEqual(i Iterator) (b bool) //是否不同
}
type BidirectionalIterator interface {
	Iterator
	HasPre() (b bool) //判断该迭代器是否可以前移
	Pre() (b bool)    //将该迭代器前移一位
}

type RandomIterator interface {
	BidirectionalIterator
	//Get(idx int) (I RandomIterator) //将该迭代器设为位于第idx节点并返回该迭代器
	GetIdx() (i int)   //获取当前的idx
	Advance(steps int) //移动steps步数
}

func Advance(it Iterator, n int) { //移动迭代器
	switch it := it.(type) {
	case RandomIterator:
		it.Advance(n)
	case BidirectionalIterator:
		if n > 0 {
			for ; n >= 0; n-- {
				it.Next()
			}
		} else if n < 0 {
			for ; n <= 0; n++ {
				it.Pre()
			}
		}
	default:
		if n > 0 {
			for ; n >= 0; n-- {
				it.Next()
			}
		}
	}
}

func Distance(first, last Iterator) (dist int) { //计算距离
	switch first := first.(type) {
	case RandomIterator:
		dist = last.(RandomIterator).GetIdx() - first.GetIdx()
	default:
		for ; first != last; first.Next() {
			dist++
		}
	}
	return
}
