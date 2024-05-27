package mystl

import (
	"fmt"
	"testing"
)

func TestVectorNew(t *testing.T) {
	v := NewVector[int](2, 1)
	for i := 0; i < 2; i++ {
		if v.At(i) != 1 {
			t.Errorf(`v.At(%v) != 1`, i)
		}
	}
}

func TestVectorNewDefault(t *testing.T) {
	v := NewVectorDefault[int]()
	if v.Size() != 0 {
		t.Error("v Not Empty!")
	}
	if v.Cap() != 1 {
		t.Error("v Cap != 1")
	}
}

func TestVectorSort(t *testing.T) {
	v := NewVectorDefault[int]()
	v.data = []int{5, 3, 1, 4, 6, 2}
	inorder := []int{1, 2, 3, 4, 5, 6}
	reverse := []int{6, 5, 4, 3, 2, 1}
	fmt.Printf("Base:%v\n", v)
	v.Sort()
	fmt.Printf("After Inorder Sort:%v\n", v)
	for i := 0; i < v.Size(); i++ {
		if v.At(i) != inorder[i] {
			t.Errorf("v[i:%v]:%v != inorder[i:%v]:%v", i, v.At(i), i, inorder[i])
		}
	}
	cmp := func(a, b int) int {
		return b - a
	}
	v.Sort(cmp)
	fmt.Printf("After Reverse Sort:%v\n", v)
	for i := 0; i < v.Size(); i++ {
		if v.At(i) != reverse[i] {
			t.Errorf("v[i:%v]:%v != reverse[i:%v]:%v", i, v.At(i), i, reverse[i])
		}
	}
}

func TestVectorClear(t *testing.T) {
	v := NewVectorDefault[int]()
	v.data = []int{5, 3, 1, 4, 6, 2}
	if v.Size() != 6 {
		t.Errorf("t.Size() Not Correct, v:%v", v)
	}
	v.Clear()
	if v.Size() != 0 {
		t.Errorf("v.Size() Not 0, v:%v", v)
	}
	if !v.Empty() {
		t.Error("v Not Empty")
	}
}

func TestVectorPushBack(t *testing.T) {
	var v *Vector[int] = NewVectorDefault[int]()
	v.PushBack(1)
	if v == nil {
		t.Error("after pushback v still nil!")
	}
	if v.Size() != 1 || v.At(0) != 1 {
		t.Errorf("v from nil to pushback 1 Not Work,v:%v", v)
	}
	v.PushBack(2)
	if v.Size() != 2 || v.At(1) != 2 {
		t.Errorf("v append not work, v:%v", v)
	}
	if v.Cap() < 2 {
		t.Errorf("v.Cap Not correct, cap:%v", v.Cap())
	}
	fmt.Printf("v.Cap after push 2: cap:%v\n", v.Cap())
}

func TestVectorPopBack(t *testing.T) {
	v := NewVectorDefault[int]()
	v.PushBack(1)
	if v.Size() != 1 {
		t.Error("v.Size != 1")
	}
	x := v.Back()
	if x != 1 {
		t.Errorf("v.Back() != 1")
	}

	v.PopBack()
	if v.Size() != 0 {
		t.Errorf("v not empty after popback:%v", v)
	}
}

func TestVectorInsert(t *testing.T) {
	var v *Vector[int]
	if err := v.Insert(0, 1); err == nil {
		t.Error("nil.Insert should be error")
	} else {
		fmt.Println(err)
	}
	v = NewVectorDefault[int]()
	if err := v.Insert(0, 1); err == nil {
		t.Error("empty insert should be error")
	} else {
		fmt.Println(err)
	}
	v.data = []int{0, 1, 2}
	if err := v.Insert(0, 3); err != nil {
		t.Error("v.insert should no error")
	} else {
		bench := []int{3, 0, 1, 2}
		if v.Size() != 4 {
			t.Errorf("v.Size wrong, size:%v", v.Size())
		}
		for i := 0; i < 4; i++ {
			if bench[i] != v.At(i) {
				t.Errorf("v differ from bench at idx:%v, v:%v, bech:%v", i, v.At(i), bench[i])
			}
		}
	}
}

func TestVectorErase(t *testing.T) {
	var v *Vector[int]
	if err := v.Erase(0); err == nil {
		t.Error("nil.Erase should be error")
	} else {
		fmt.Println(err)
	}
	v = NewVectorDefault[int]()
	if err := v.Erase(0); err == nil {
		t.Error("empty insert should be error")
	} else {
		fmt.Println(err)
	}
	v.data = []int{0, 1, 2}
	if err := v.Erase(0); err != nil {
		t.Error("v.insert should no error")
	} else {
		bench := []int{1, 2}
		if v.Size() != 2 {
			t.Errorf("v.Size wrong, size:%v", v.Size())
		}
		for i := 0; i < 2; i++ {
			if bench[i] != v.At(i) {
				t.Errorf("v differ from bench at idx:%v, v:%v, bech:%v", i, v.At(i), bench[i])
			}
		}
	}
}

func TestVectorReverse(t *testing.T) {
	v := NewVectorDefault[int]()
	v.data = []int{1, 2, 3, 4}
	bench := []int{4, 3, 2, 1}
	v.Reverse()
	if v.Size() != 4 {
		t.Errorf("Wrong Size:%d", v.Size())
	}
	for i := 0; i < 4; i++ {
		if bench[i] != v.At(i) {
			t.Errorf("benc differs from v, i:%d, bench:%v, v:%v", i, bench[i], v.At(i))
		}
	}
}

func TestVectorData(t *testing.T) {
	//v := NewDefault[int]()
	v := NewVectorData[int]([]int{1, 2, 3, 4})
	for i, e := range v.data {
		if e != v.At(i) {
			t.Errorf("wrong at idx:%d, e:%v", i, e)
		}
	}
}
