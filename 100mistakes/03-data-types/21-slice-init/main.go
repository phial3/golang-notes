package main

func convertEmptySlice(foos []Foo) []Bar {
	bars := make([]Bar, 0)

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenCapacity(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenLength(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))

	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}

type Foo struct{}

type Bar struct{}

func fooToBar(foo Foo) Bar {
	return Bar{}
}
