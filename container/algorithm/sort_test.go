package algorithm

import (
	"github.com/phial3/container"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	Sort(arr, container.ASC[int])

	if !IsSorted(arr, container.ASC[int]) {
		t.Error("not sorted!")
	}
}

func TestBinarySearch(t *testing.T) {
	var count = 1000
	var arr = make([]int, 0, count)
	for i := 0; i < count; i++ {
		arr = append(arr, rand.Intn(count))
	}
	Unique(&arr, func(x int) int { return x })
	Sort(arr, container.DESC[int])

	var m = make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}

	for k, v := range m {
		res := BinarySearch(arr, k, container.DESC[int])
		if res != v {
			t.Fatal("error!")
		}
	}

	if res := BinarySearch(arr, count, container.DESC[int]); res != -1 {
		t.Fatal("error!")
	}
}
