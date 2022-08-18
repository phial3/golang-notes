package slices_test

import (
	"testing"

	"github.com/phial3/genesis/slices"
	"github.com/matryer/is"
)

func TestConcat(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected []int) {
		is.Equal(slices.Concat(given...), expected)
	}
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
	f([][]int{{1}}, []int{1})
	f([][]int{{1}, {}}, []int{1})
	f([][]int{{}, {1}}, []int{1})
	f([][]int{{1, 2}, {3, 4, 5}}, []int{1, 2, 3, 4, 5})
}

func TestProduct2(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Product2(given...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int{{1, 2}, {3}, {4, 5}}, [][]int{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestZip(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Zip(given...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f([][]int{}, [][]int{})
	f([][]int{{1}, {2}}, [][]int{{1, 2}})
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}})
	f([][]int{{1, 2, 3}, {4, 5}}, [][]int{{1, 4}, {2, 5}})
}
