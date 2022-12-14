package slice

import (
	"github.com/phial3/container"
	"github.com/phial3/container/algorithm"
	"github.com/phial3/container/internal/utils"
	"testing"
)

func TestSlice_Sort(t *testing.T) {
	var arr Slice[int] = New[int]()
	arr.Push(1, 3, 5, 7, 9, 2, 4, 6, 8, 0)
	arr.Sort(dao.ASC[int])
	if !algorithm.IsSorted(arr, dao.ASC[int]) {
		t.Fatal("error!")
	}
}

func TestSlice_Unique(t *testing.T) {
	var n = 1000
	var arr1 = New[int]()
	for i := 0; i < n; i++ {
		arr1.Push(utils.Rand.Intn(100))
	}
	arr1.Unique(dao.ASC[int])

	var length = arr1.Len()
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr1[i] == arr1[j] {
				t.Fatal("error!")
			}
		}
	}
}

func TestSlice_New(t *testing.T) {
	var a = New[int]()
	a.Push(1, 3, 5, 7)
	a.ForEach(func(iter *Iterator[int]) {
		println(iter.Index, iter.Value)
		if iter.Value == 5 {
			iter.Break()
		}
	})
}
