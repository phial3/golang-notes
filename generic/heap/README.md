<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# heap

```go
import "github.com/phial3/generic/heap"
```

Package heap provides an implementation of a binary heap\. A binary heap \(binary min\-heap\) is a tree with the property that each node is the minimum\-valued node in its subtree\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/phial3/generic/heap"
)

func main() {
	heap := heap.New(func(a, b int) bool { return a < b })

	heap.Push(5)
	heap.Push(2)
	heap.Push(3)

	v, _ := heap.Pop()
	fmt.Println(v)

	v, _ = heap.Peek()
	fmt.Println(v)
}
```

#### Output

```
2
3
```

</p>
</details>

## Index

- [type Heap](<#type-heap>)
  - [func From[T any](less g.LessFn[T], t ...T) *Heap[T]](<#func-from>)
  - [func FromSlice[T any](less g.LessFn[T], data []T) *Heap[T]](<#func-fromslice>)
  - [func New[T any](less g.LessFn[T]) *Heap[T]](<#func-new>)
  - [func (h *Heap[T]) Peek() (T, bool)](<#func-heapt-peek>)
  - [func (h *Heap[T]) Pop() (T, bool)](<#func-heapt-pop>)
  - [func (h *Heap[T]) Push(x T)](<#func-heapt-push>)
  - [func (h *Heap[T]) Size() int](<#func-heapt-size>)


## type [Heap](<https://github.com/phial3/generic/blob/master/heap/heap.go#L11-L14>)

Heap implements a binary heap\.

```go
type Heap[T any] struct {
    // contains filtered or unexported fields
}
```

### func [From](<https://github.com/phial3/generic/blob/master/heap/heap.go#L25>)

```go
func From[T any](less g.LessFn[T], t ...T) *Heap[T]
```

From returns a new heap with the given less function and initial data\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/phial3/generic/heap"
)

func main() {
	heap := heap.From(func(a, b int) bool { return a < b }, 5, 2, 3)

	v, _ := heap.Pop()
	fmt.Println(v)

	v, _ = heap.Peek()
	fmt.Println(v)
}
```

#### Output

```
2
3
```

</p>
</details>

### func [FromSlice](<https://github.com/phial3/generic/blob/master/heap/heap.go#L31>)

```go
func FromSlice[T any](less g.LessFn[T], data []T) *Heap[T]
```

FromSlice returns a new heap with the given less function and initial data\. The \`data\` is not copied and used as the inside array\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/phial3/generic/heap"
)

func main() {
	heap := heap.FromSlice(func(a, b int) bool { return a > b }, []int{-1, 5, 2, 3})

	v, _ := heap.Pop()
	fmt.Println(v)

	v, _ = heap.Peek()
	fmt.Println(v)
}
```

#### Output

```
5
3
```

</p>
</details>

### func [New](<https://github.com/phial3/generic/blob/master/heap/heap.go#L17>)

```go
func New[T any](less g.LessFn[T]) *Heap[T]
```

New returns a new heap with the given less function\.

### func \(\*Heap\[T\]\) [Peek](<https://github.com/phial3/generic/blob/master/heap/heap.go#L68>)

```go
func (h *Heap[T]) Peek() (T, bool)
```

Peek returns the minimum element from the heap without removing it\. if the heap is empty\, it returns zero value and false\.

### func \(\*Heap\[T\]\) [Pop](<https://github.com/phial3/generic/blob/master/heap/heap.go#L51>)

```go
func (h *Heap[T]) Pop() (T, bool)
```

Pop removes and returns the minimum element from the heap\. If the heap is empty\, it returns zero value and false\.

### func \(\*Heap\[T\]\) [Push](<https://github.com/phial3/generic/blob/master/heap/heap.go#L44>)

```go
func (h *Heap[T]) Push(x T)
```

Push pushes the given element onto the heap\.

### func \(\*Heap\[T\]\) [Size](<https://github.com/phial3/generic/blob/master/heap/heap.go#L78>)

```go
func (h *Heap[T]) Size() int
```

Size returns the number of elements in the heap\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)