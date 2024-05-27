package mystl

import "fmt"

type matrixer[T MyArithmetic] interface {
	Set(m int, n int, val []T)                         // val size is m*n
	Assign(b matrixer[T])                              //仅修改切片指针
	Copy(b matrixer[T])                                //拷贝数据
	Equal(b matrixer[T]) bool                          //是否相同
	Add(b matrixer[T]) (c matrixer[T], err error)      //矩阵加法
	Multiply(b matrixer[T]) (c matrixer[T], err error) //矩阵乘法
	Transpose() (oth matrixer[T])                      //转置
	Row(idx int) *[]T                                  //获取行指针
	Cell(x int, y int) *T                              //获取元素指针
	At(x int, y int) T                                 //获取值
	Data() *[][]T                                      //获取全部数据的指针
	RowNum() int                                       //获取行数
	ColNum() int                                       //获取列数
	Shape() (row, col int)                             //获取矩阵形状
	IsSquare() bool                                    //是否是方阵
}

func NewMatrix[T MyArithmetic](m int, n int, init T) *Matrix[T] {
	if m < 0 || n < 0 {
		fmt.Printf("param invalid, m:%v,n:%v\n", m, n)
		return nil
	}
	mat := &Matrix[T]{data: make([][]T, m)}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat.data[i] = append(mat.data[i], init)
		}
	}
	return mat
}

func NewMatrixVector[T MyArithmetic](m int, n int, val []T) *Matrix[T] {
	if m < 0 || n < 0 {
		fmt.Printf("param invalid, m:%v,n:%v\n", m, n)
		return nil
	}
	mat := &Matrix[T]{data: make([][]T, m)}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat.data[i] = append(mat.data[i], val[i*n+j])
		}
	}
	return mat
}

type Matrix[T MyArithmetic] struct {
	data [][]T
}

func (mat *Matrix[T]) Set(m int, n int, val []T) { // val size is m*n
	if mat == nil {
		panic("mat is nil")
	}
	if m < 0 || n < 0 {
		fmt.Printf("param invalid, m:%v,n:%v\n", m, n)
		return
	}
	if len(val) != m*n {
		fmt.Printf("Error, len(val):%v differs from m*n:%v\n", len(val), m*n)
		return
	}
	mat.data = make([][]T, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat.data[i] = append(mat.data[i], val[i*n+j])
		}
	}
}

func (mat *Matrix[T]) Assign(b matrixer[T]) { //仅修改切片指针, 浅拷贝
	if mat == nil {
		panic("mat is nil")
	}
	mat.data = *b.Data()
}

func (mat *Matrix[T]) Copy(b matrixer[T]) { //拷贝数据
	if mat == nil {
		panic("mat is nil")
	}
	bdata := b.Data()
	m, n := b.Shape()
	mat.data = make([][]T, b.RowNum())
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat.data[i] = append(mat.data[i], (*bdata)[i][j])
		}
	}
}

func (mat *Matrix[T]) Equal(b matrixer[T]) bool { //是否相同
	if mat == nil {
		panic("mat is nil")
	}
	mr, mc := mat.Shape()
	br, bc := b.Shape()

	if mr != br || mc != bc {
		fmt.Printf("Debug, this.Shape(%v,%v) != other.Shape(%v,%v)", mr, mc, br, bc)
		return false
	}

	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if mat.At(i, j) != b.At(i, j) {
				fmt.Printf("Debug, At(%d,%d), this:%v, other:%v\n", i, j, mat.At(i, j), b.At(i, j))
				return false
			}
		}
	}
	return true
}

func (mat *Matrix[T]) Add(b matrixer[T]) (c matrixer[T], err error) { //矩阵加法
	if mat == nil {
		panic("mat is nil")
	}
	mr, mc := mat.Shape()
	br, bc := b.Shape()

	if mr != br || mc != bc {
		return nil, fmt.Errorf("this.Shape(%v,%v) != other.Shape(%v,%v)", mr, mc, br, bc)
	}
	var d T
	res := NewMatrix[T](0, 0, d)
	res.Copy(mat)
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			res.data[i][j] += b.At(i, j)
		}
	}
	return res, nil
}
func (mat *Matrix[T]) Multiply(b matrixer[T]) (c matrixer[T], err error) { //矩阵乘法
	if mat == nil {
		panic("mat is nil")
	}
	mr, mc := mat.Shape()
	br, bc := b.Shape()
	if mc != br {
		return nil, fmt.Errorf("cannot Multiply this(%v,%v) with other(%v,%v)", mr, mc, br, bc)
	}
	var d T
	res := NewMatrix(mr, bc, d)
	for i := 0; i < mr; i++ {
		for j := 0; j < bc; j++ {
			for k := 0; k < mc; k++ {
				res.data[i][j] += mat.At(i, k) * b.At(k, j)
			}
		}
	}
	return res, nil
}

func (mat *Matrix[T]) Transpose() (oth matrixer[T]) { //转置
	if mat == nil {
		panic("mat is nil")
	}
	row, col := mat.Shape()
	var d T
	oth = NewMatrix[T](col, row, d)
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			*oth.Cell(i, j) = mat.At(j, i)
		}
	}
	return
}

func (mat *Matrix[T]) Row(idx int) *[]T { //获取行指针
	if mat == nil {
		panic("mat is nil")
	}
	return &mat.data[idx]
}

func (mat *Matrix[T]) Cell(x int, y int) *T { //获取元素指针
	if mat == nil {
		panic("mat is nil")
	}
	return &mat.data[x][y]
}
func (mat *Matrix[T]) At(x int, y int) T { //获取值
	if mat == nil {
		panic("mat is nil")
	}
	return mat.data[x][y]
}
func (mat *Matrix[T]) Data() *[][]T { //获取全部数据的指针
	if mat == nil {
		panic("mat is nil")
	}
	return &mat.data
}
func (mat *Matrix[T]) RowNum() int { //获取行数
	if mat == nil {
		panic("mat is nil")
	}
	return len(mat.data)
}
func (mat *Matrix[T]) ColNum() int { //获取列数
	if mat == nil {
		panic("mat is nil")
	}
	return len(mat.data[0])
}
func (mat *Matrix[T]) Shape() (row, col int) { //获取矩阵形状
	if mat == nil {
		panic("mat is nil")
	}
	return mat.RowNum(), mat.ColNum()
}
func (mat *Matrix[T]) IsSquare() bool { //是否是方阵
	if mat == nil {
		panic("mat is nil")
	}
	return mat.RowNum() == mat.ColNum()
}
