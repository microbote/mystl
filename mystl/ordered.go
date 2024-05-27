package mystl

import "cmp"

type MyOrdered interface {
	cmp.Ordered
}

type MyArithmetic interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Load interface {
	any
}

type Comparator[T MyOrdered] func(a, b T) int

type ComparatorAny[T Load] func(a, b T) int
