package benchmark

import (
	"github.com/phial3/container/internal/utils"
	"github.com/phial3/container/rapid"
	"testing"
)

type entry struct {
	Key string
	Val int
}

//func (c entry) Equal(x *entry) bool {
//	return c.Key == x.Key
//}

func BenchmarkRapid_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rapid.New(bench_count, func(a, b *entry) bool {
			return a.Key == b.Key
		})
	}
}

func BenchmarkRapid_Append(b *testing.B) {
	var arr = make([]string, 0, bench_count)
	var val = 1
	for i := 0; i < bench_count; i++ {
		arr = append(arr, utils.Alphabet.Generate(8))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var r = rapid.New(bench_count, func(a, b *entry) bool {
			return a.Key == b.Key
		})
		var id1 = r.NextID()
		var q1 = rapid.EntryPoint{Head: id1, Tail: id1}
		var id2 = r.NextID()
		var q2 = rapid.EntryPoint{Head: id2, Tail: id2}

		for j := 0; j < bench_count/2; j++ {
			r.Append(&q1, &entry{Key: arr[j], Val: val})
		}
		for j := 0; j < bench_count/2; j++ {
			r.Append(&q2, &entry{Key: arr[bench_count/2+j], Val: val})
		}
	}
	b.StopTimer()
}
