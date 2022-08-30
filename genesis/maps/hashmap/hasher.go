package hashmap

import (
	"hash/fnv"
	"unsafe"
)

type Hasher interface {
	Hash() uint32
	Equals(a any) bool
}

type HashFunc[K any] func(k K) uint32

type EqualsFunc[K any] func(k1, k2 K) bool

func hasherHashFunc[K Hasher]() HashFunc[K] {
	return func(k K) uint32 {
		return k.Hash()
	}
}

func hasherEqualsFunc[K Hasher]() EqualsFunc[K] {
	return func(k1, k2 K) bool {
		return k1.Equals(k2)
	}
}

func comparableHashFunc[K comparable]() HashFunc[K] {
	return func(k K) uint32 {
		b := *(*[]byte)(unsafe.Pointer(&struct {
			data unsafe.Pointer
			len  int
		}{unsafe.Pointer(&k), int(unsafe.Sizeof(k))}))
		h := fnv.New32a()
		_, _ = h.Write(b)
		return h.Sum32()
	}
}

func comparableEqualsFunc[K comparable]() EqualsFunc[K] {
	return func(k1, k2 K) bool {
		return k1 == k2
	}
}

var stringHashFunc HashFunc[string] = func(k string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(k))
	return h.Sum32()
}

var stringEqualsFunc EqualsFunc[string] = func(k1, k2 string) bool {
	return k1 == k2
}

var intHashFunc HashFunc[int] = func(k int) uint32 {
	return uint32(k)
}

var intEqualsFunc EqualsFunc[int] = func(k1, k2 int) bool {
	return k1 == k2
}
