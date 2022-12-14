package benchmark

import (
	"github.com/phial3/container/segment_tree"
	"math/rand"
	"testing"
)

func BenchmarkSegmentTree_Query(b *testing.B) {
	var arr = make([]int, 0)
	for i := 0; i < bench_count; i++ {
		arr = append(arr, testvals[i])
	}
	var tree = segment_tree.New(arr, segment_tree.Init[int], segment_tree.Merge[int])

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var left = rand.Intn(bench_count)
			var right = rand.Intn(bench_count)
			if left > right {
				left, right = right, left
			}
			tree.Query(left, right)
		}
	}
}

func BenchmarkSegmentTree_Update(b *testing.B) {
	var arr1 = make([]int, 0)
	for i := 0; i < bench_count; i++ {
		arr1 = append(arr1, testvals[i])
	}
	var tree = segment_tree.New(arr1, segment_tree.Init[int], segment_tree.Merge[int])

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var x = rand.Intn(bench_count)
			tree.Update(x, x)
		}
	}
}
