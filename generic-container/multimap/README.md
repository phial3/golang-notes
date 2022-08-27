<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# multimap

```go
import "github.com/phial3/generic/multimap"
```

Package multimap provides an associative container that permits multiple entries with the same key\.

There are four implementations of the MultiMap data structure\, identified by separate New\* functions\. They differ in the following ways: \- whether key type and value type must be comparable\. \- whether duplicate entries \(same key and same value\) are permitted\. \- whether keys and values are sorted or unsorted in Get\, Each\, and EachAssociation methods\.

## Index

- [type MultiMap](<#type-multimap>)
  - [func NewAvlSet[K, V any](keyLess g.LessFn[K], valueLess g.LessFn[V]) MultiMap[K, V]](<#func-newavlset>)
  - [func NewAvlSlice[K any, V comparable](keyLess g.LessFn[K]) MultiMap[K, V]](<#func-newavlslice>)
  - [func NewMapSet[K comparable, V any](valueLess g.LessFn[V]) MultiMap[K, V]](<#func-newmapset>)
  - [func NewMapSlice[K, V comparable]() MultiMap[K, V]](<#func-newmapslice>)


## type [MultiMap](<https://github.com/phial3/generic/blob/master/multimap/multimap.go#L11-L40>)

MultiMap is an associative container that contains a list of key\-value pairs\, while permitting multiple entries with the same key\.

```go
type MultiMap[K, V any] interface {
    // Dimension returns number of distinct keys.
    Dimension() int
    // Size returns total number of entries.
    Size() int

    // Count returns number of entries with a given key.
    Count(key K) int
    // Has determines whether at least one entry exists with a given key.
    Has(key K) bool
    // Get returns a list of values with a given key.
    Get(key K) []V

    // Put adds an entry.
    // Whether duplicate entries are allowed depends on the chosen implementation.
    Put(key K, value V)
    // Remove removes an entry.
    // If duplicate entries are allowed, this removes only one entry.
    // This is a no-op if the entry does not exist.
    Remove(key K, value V)
    // RemoveAll removes every entry with a given key.
    RemoveAll(key K)
    // Clear deletes all entries.
    Clear()

    // Each calls 'fn' on every entry.
    Each(fn func(key K, value V))
    // EachAssociation calls 'fn' on every key and list of values.
    EachAssociation(fn func(key K, values []V))
}
```

### func [NewAvlSet](<https://github.com/phial3/generic/blob/master/multimap/avl.go#L109>)

```go
func NewAvlSet[K, V any](keyLess g.LessFn[K], valueLess g.LessFn[V]) MultiMap[K, V]
```

NewAvlSet creates a MultiMap using AVL tree and AVL set\. \- Duplicate entries are not permitted\. \- Both keys and values are sorted\.

### func [NewAvlSlice](<https://github.com/phial3/generic/blob/master/multimap/avl.go#L95>)

```go
func NewAvlSlice[K any, V comparable](keyLess g.LessFn[K]) MultiMap[K, V]
```

NewAvlSlice creates a MultiMap using AVL tree and builtin slice\. \- Value type must be comparable\. \- Duplicate entries are permitted\. \- Keys are sorted\, but values are unsorted\.

### func [NewMapSet](<https://github.com/phial3/generic/blob/master/multimap/map.go#L108>)

```go
func NewMapSet[K comparable, V any](valueLess g.LessFn[V]) MultiMap[K, V]
```

NewMapSet creates a MultiMap using builtin map and AVL set\. \- Key type must be comparable\. \- Duplicate entries are not permitted\. \- Values are sorted\, but keys are unsorted\.

### func [NewMapSlice](<https://github.com/phial3/generic/blob/master/multimap/map.go#L94>)

```go
func NewMapSlice[K, V comparable]() MultiMap[K, V]
```

NewMapSlice creates a MultiMap using builtin map and builtin slice\. \- Both key type and value type must be comparable\. \- Duplicate entries are permitted\. \- Both keys and values are unsorted\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)