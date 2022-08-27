package segment_tree

import (
	"github.com/phial3/container"
	"github.com/phial3/container/algorithm"
)

type Schema[T container.Number[T]] struct {
	MaxValue T
	MinValue T
	Sum      T
}

func Init[T container.Number[T]](op Operate, x T) Schema[T] {
	var result = Schema[T]{
		MaxValue: x,
		MinValue: x,
		Sum:      x,
	}
	if op == Operate_Query {
		result.Sum = 0
	}
	return result
}

func Merge[T container.Number[T]](a, b Schema[T]) Schema[T] {
	return Schema[T]{
		MaxValue: algorithm.Max(a.MaxValue, b.MaxValue),
		MinValue: algorithm.Min(a.MinValue, b.MinValue),
		Sum:      a.Sum + b.Sum,
	}
}
