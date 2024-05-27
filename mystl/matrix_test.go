package mystl

import (
	"testing"
)

func TestMatrixData(t *testing.T) {
	mat := NewMatrix[int](5, 4, 1)
	//fmt.Println(mat)
	row, col := mat.Shape()
	if row != 5 {
		t.Errorf("row(%v) != 5", row)
		return
	}
	if col != 4 {
		t.Errorf("col(%v) != 4", col)
		return
	}
	if mat.IsSquare() {
		t.Error("should not be square")
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if mat.At(i, j) != 1 {
				t.Errorf("mat[%d][%d]: %v not correct", i, j, mat.At(i, j))
			}
		}
	}
	*mat.Cell(0, 0) = 55
	if mat.At(0, 0) != 55 {
		t.Errorf("Cell not work")
	}
	refRow := mat.Row(1)
	for j := 0; j < 4; j++ {
		(*refRow)[j] = 3
	}
	for j := 0; j < 4; j++ {
		if mat.At(1, j) != 3 {
			t.Errorf("mat[1][%d] != 3", j)
			break
		}
	}
}

func TestMatrixSet(t *testing.T) {
	val := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mat := NewMatrix[int](0, 0, 0)
	mat.Set(4, 2, val)
	row, col := mat.Shape()
	if row != 4 || col != 2 {
		t.Errorf("row:%d,col:%d not match", row, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat.At(i, j) != val[i*col+j] {
				t.Errorf("set error, val not match, i:%d,j:%d, mat:%v, val:%v", i, j, mat.At(i, j), val[i*row+j])
			}
		}
	}

	c := NewMatrix(0, 0, 0)
	c.Assign(mat)
	crow, ccol := c.Shape()
	if crow != row || ccol != col {
		t.Errorf("Assign failed, crow:%d,ccol:%d,row:%d,col:%d", crow, ccol, row, col)
	}
	if !c.Equal(mat) {
		t.Errorf("Equal Error")
	}

	c.data[0][0] = 99
	if !c.Equal(mat) {
		t.Errorf("Assign Error, Because c is referenced by mat")
	}

	d := NewMatrix(0, 0, 0)
	d.Copy(mat)
	if !d.Equal(mat) {
		t.Errorf("copy error")
	}
	d.data[0][0] = 71
	if d.Equal(mat) {
		t.Errorf("copy error, because d is deep copy of mat")
	}
}

func TestMatrixAdd(t *testing.T) {
	a := NewMatrix(0, 0, 0)
	a.Set(4, 2, []int{1, 2, 3, 4, 5, 6, 7, 8})
	b := NewMatrix(4, 1, 1)
	if _, err := a.Add(b); err == nil {
		t.Errorf("Shape Dismatch should be error")
	}
	c := NewMatrix(4, 2, 1)
	if e, err := a.Add(c); err != nil {
		t.Errorf("add error,%q", err)
	} else {
		d := NewMatrixVector(4, 2, []int{2, 3, 4, 5, 6, 7, 8, 9})
		if !e.Equal(d) {
			t.Errorf("add not correct, result:%v", e)
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	a := NewMatrixVector(4, 2, []int{1, 2, 3, 4, 5, 6, 7, 8})
	b := NewMatrixVector(2, 1, []int{1, 0})
	c := NewMatrixVector(4, 1, []int{1, 3, 5, 7})
	if _, err := a.Multiply(c); err == nil {
		t.Errorf("shape dismatch should be err")
	}
	d, err := a.Multiply(b)
	if err != nil {
		t.Errorf("Multiply failed, %q", err)
	} else {
		if !d.Equal(c) {
			t.Errorf("Multiply result error, %v", d)
		}
	}
}

func TestMatrixTransponse(t *testing.T) {
	a := NewMatrixVector(4, 2, []int{1, 2, 3, 4, 5, 6, 7, 8})
	/*
		1 2  -> 1 3 5 7
		3 4		2 4 6 8
		5 6
		7 8
	*/
	b := NewMatrixVector(2, 4, []int{1, 3, 5, 7, 2, 4, 6, 8})
	c := a.Transpose()
	if !c.Equal(b) {
		t.Errorf("transpose error, %v", c)
	}
}
