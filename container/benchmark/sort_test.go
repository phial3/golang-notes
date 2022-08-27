package benchmark

import (
	"github.com/phial3/container"
	"github.com/phial3/container/algorithm"
	"sort"
	"testing"
)

func BenchmarkSort_Quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count, bench_count)
		copy(arr, testvals[:bench_count])
		algorithm.Sort(arr, dao.ASC[int])
	}
}

func BenchmarkSort_Golang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count, bench_count)
		copy(arr, testvals[:bench_count])
		sort.Ints(arr)
	}
}
