package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/phial3/automi/collectors"
	"github.com/phial3/automi/stream"
)

func main() {
	data := []string{
		"12528 52 2 0.04 50 0.96 0 0 52 100 0 0 0 0 0 0 0 0 46 0.88 0 0 6 0.12 0 0 52 100 1 0.02 51 0.98 0 0 0 0 52 100 21 0.4 31 0.6 0 0 52 100",
		"12701 87 49 0.56 38 0.44 0 0 87 100 0 0 0 0 0 0 0 0 83 0.95 0 0 2 0.02 2 0.02 87 99 0 0 87 1 0 0 0 0 87 100 18 0.21 69 0.79 0 0 87 100",
		"12733 13 6 0.46 7 0.54 0 0 13 100 0 0 0 0 0 0 0 0 13 1 0 0 0 0 0 0 13 100 0 0 13 1 0 0 0 0 13 100 2 0.15 11 0.85 0 0 13 100",
		"12734 252 170 0.67 82 0.33 0 0 252 100 0 0 0 0 0 0 0 0 243 0.96 0 0 4 0.02 5 0.02 252 100 0 0 252 1 0 0 0 0 252 100 61 0.24 191 0.76 0 0 252 100",
		"12737 35 34 0.97 1 0.03 0 0 35 100 0 0 0 0 0 0 0 0 35 1 0 0 0 0 0 0 35 100 2 0.06 33 0.94 0 0 0 0 35 100 20 0.57 15 0.43 0 0 35 100",
		"12750 37 36 0.97 1 0.03 0 0 37 100 0 0 0 0 0 0 0 0 37 1 0 0 0 0 0 0 37 100 0 0 37 1 0 0 0 0 37 100 21 0.57 16 0.43 0 0 37 100",
		"12751 7 2 0.29 5 0.71 0 0 7 100 0 0 0 0 0 0 0 0 7 1 0 0 0 0 0 0 7 100 0 0 7 1 0 0 0 0 7 100 2 0.29 5 0.71 0 0 7 100",
		"12754 134 64 0.48 70 0.52 0 0 134 100 2 0.01 0 0 0 0 0 0 127 0.95 0 0 4 0.03 1 0.01 134 100 1 0.01 133 0.99 0 0 0 0 134 100 27 0.2 107 0.8 0 0 134 100",
		"12758 82 50 0.61 32 0.39 0 0 82 100 0 0 0 0 0 0 0 0 76 0.93 0 0 6 0.07 0 0 82 100 0 0 82 1 0 0 0 0 82 100 24 0.29 58 0.71 0 0 82 100",
		"12759 11 0 0 11 1 0 0 11 100 0 0 0 0 0 0 0 0 11 1 0 0 0 0 0 0 11 100 0 0 11 1 0 0 0 0 11 100 8 0.73 3 0.27 0 0 11 100",
		"12763 11 11 1 0 0 0 0 11 100 0 0 0 0 0 0 0 0 11 1 0 0 0 0 0 0 11 100 0 0 11 1 0 0 0 0 11 100 8 0.73 3 0.27 0 0 11 100",
		"12764 64 61 0.95 3 0.05 0 0 64 100 0 0 0 0 0 0 0 0 62 0.97 0 0 2 0.03 0 0 64 100 0 0 64 1 0 0 0 0 64 100 5 0.08 59 0.92 0 0 64 100",
		"12768 63 1 0.02 62 0.98 0 0 63 100 0 0 0 0 0 0 0 0 61 0.97 0 0 1 0.02 1 0.02 63 100 0 0 63 1 0 0 0 0 63 100 23 0.37 40 0.63 0 0 63 100",
		"12779 242 183 0.76 59 0.24 0 0 242 100 0 0 0 0 0 0 0 0 239 0.99 0 0 3 0.01 0 0 242 100 1 0 241 1 0 0 0 0 242 100 155 0.64 87 0.36 0 0 242 100",
		"12783 201 66 0.33 135 0.67 0 0 201 100 0 0 0 0 2 0.01 0 0 195 0.97 0 0 3 0.01 1 0 201 99 4 0.02 197 0.98 0 0 0 0 201 100 77 0.38 124 0.62 0 0 201 100",
		"12786 4 3 0.75 1 0.25 0 0 4 100 0 0 0 0 0 0 0 0 4 1 0 0 0 0 0 0 4 100 0 0 4 1 0 0 0 0 4 100 1 0.25 3 0.75 0 0 4 100",
		"12788 83 39 0.47 44 0.53 0 0 83 100 0 0 0 0 0 0 0 0 81 0.98 0 0 2 0.02 0 0 83 100 0 0 83 1 0 0 0 0 83 100 35 0.42 48 0.58 0 0 83 100",
		"12789 272 115 0.42 157 0.58 0 0 272 100 0 0 0 0 0 0 0 0 262 0.96 0 0 6 0.02 4 0.01 272 99 1 0 271 1 0 0 0 0 272 100 70 0.26 202 0.74 0 0 272 100",
		"13731 17 2 0.12 15 0.88 0 0 17 100 0 0 0 0 0 0 0 0 15 0.88 0 0 2 0.12 0 0 17 100 0 0 17 1 0 0 0 0 17 100 7 0.41 10 0.59 0 0 17 100",
	}

	var strBuilder = bytes.NewBufferString("")
	csvSink := collectors.CSV(strBuilder)
	csvSink.DelimChar('|')

	stream := stream.New(data)

	stream.Map(func(row string) []string {
		return strings.Split(row, " ")
	})

	stream.Into(csvSink)

	if err := <-stream.Open(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strBuilder.String())
}
