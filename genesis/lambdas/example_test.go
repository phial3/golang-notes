package lambdas_test

import (
	"fmt"

	"github.com/phial3/genesis/lambdas"
	"github.com/phial3/genesis/slices"
)

func ExampleMust() {
	res := lambdas.Must(slices.Min([]int{42, 7, 13}))
	fmt.Println(res)
	//Output:
	// 7
}

func ExampleDefaultTo() {
	res := lambdas.DefaultTo(13)(slices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleSafe() {
	res := lambdas.Safe(slices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 0
}

func ExampleAbs() {
	res := lambdas.Abs(-13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMin() {
	res := lambdas.Min(15, 13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMax() {
	res := lambdas.Max(10, 13)
	fmt.Println(res)
	//Output:
	// 13
}
