// Package hashmap concurrent hash map
package hashmap

import (
	"errors"
	"sync"
)

const (
	defaultCapacity  = 16
	treeifyThreshold = 16
)

type node[K, V any] struct {
	hash  uint32
	key   K
	value V
	right *node[K, V]
	left  *node[K, V]
}

type bucket[K, V any] struct {
	sync.RWMutex
	node *node[K, V]
	tree bool
	size int64
}

type ConcurrentHashMap[K, V any] struct {
	capacity uint32
	table    []*bucket[K, V]
	hf       HashFunc[K]
	ef       EqualsFunc[K]
}

func NewHashMap[K, V any](hf HashFunc[K], ef EqualsFunc[K]) ConcurrentHashMap[K, V] {
	hashmap, _ := NewWithCap[K, V](defaultCapacity, hf, ef)
	return hashmap
}

func NewWithCap[K, V any](capacity int, hf HashFunc[K], ef EqualsFunc[K]) (ConcurrentHashMap[K, V], error) {
	var hashmap ConcurrentHashMap[K, V]
	if capacity <= 0 {
		err := errors.New("capacity must be positive value")
		return hashmap, err
	}
	if hf == nil || ef == nil {
		err := errors.New("hash and equals funcs cannot be nil")
		return hashmap, err
	}
	hashmap.capacity = uint32(capacity)
	hashmap.table = make([]*bucket[K, V], hashmap.capacity)
	for i := 0; i < int(hashmap.capacity); i++ {
		hashmap.table[i] = &bucket[K, V]{}
	}
	hashmap.hf = hf
	hashmap.ef = ef
	return hashmap, nil
}

func NewMap[K Hasher, V any]() ConcurrentHashMap[K, V] {
	hashmap, _ := NewMapWithCap[K, V](defaultCapacity)
	return hashmap
}

func NewMapWithCap[K Hasher, V any](capacity int) (ConcurrentHashMap[K, V], error) {
	return NewWithCap[K, V](capacity, hasherHashFunc[K](), hasherEqualsFunc[K]())
}

func NewComparableMap[K comparable, V any]() ConcurrentHashMap[K, V] {
	hashmap, _ := NewComparableMapWithCap[K, V](defaultCapacity)
	return hashmap
}

func NewComparableMapWithCap[K comparable, V any](capacity int) (ConcurrentHashMap[K, V], error) {
	return NewWithCap[K, V](capacity, comparableHashFunc[K](), comparableEqualsFunc[K]())
}

func NewStringMap[V any]() ConcurrentHashMap[string, V] {
	hashmap, _ := NewStringMapWithCap[V](defaultCapacity)
	return hashmap
}

func NewStringMapWithCap[V any](capacity int) (ConcurrentHashMap[string, V], error) {
	return NewWithCap[string, V](capacity, stringHashFunc, stringEqualsFunc)
}

func NewIntMap[V any]() ConcurrentHashMap[int, V] {
	hashmap, _ := NewIntMapWithCap[V](defaultCapacity)
	return hashmap
}

func NewIntMapWithCap[V any](capacity int) (ConcurrentHashMap[int, V], error) {
	return NewWithCap[int, V](capacity, intHashFunc, intEqualsFunc)
}

// Put maps the given key to the value, and saves the entry.
// In case of there is already an entry mapped by the given key, it updates the value of the entry.
func (m *ConcurrentHashMap[K, V]) Put(key K, val V) {
	h := m.hf(key)
	b := m.table[h%m.capacity]
	b.Lock()
	b.put(h, key, val, m.ef)
	b.Unlock()
}

// Get returns value of the entry mapped by given key.
// If there is mopping by given key, it returns false.
func (m *ConcurrentHashMap[K, V]) Get(key K) (V, bool) {
	h := m.hf(key)
	b := m.table[h%m.capacity]
	b.RLock()
	n := b.get(h, key, m.ef)
	b.RUnlock()
	if n == nil {
		return *new(V), false
	}
	return n.value, true
}

// GetOrDefault returns the value of the entry mapped by the given key.
// If there is mopping by the given key, it returns default value argument.
func (m *ConcurrentHashMap[K, V]) GetOrDefault(key K, defVal V) V {
	h := m.hf(key)
	b := m.table[h%m.capacity]
	b.RLock()
	n := b.get(h, key, m.ef)
	b.RUnlock()
	if n == nil {
		return defVal
	}
	return n.value
}

// Contains returns if there is an entry mapped by the given key.
func (m *ConcurrentHashMap[K, V]) Contains(key K) bool {
	h := m.hf(key)
	b := m.table[h%m.capacity]
	b.RLock()
	n := b.get(h, key, m.ef)
	b.RUnlock()
	return n != nil
}

// Remove removes the entry mapped by the given key and returns value of removed entry and true.
// In case of there is entry by the given key, It returns nil and false.
func (m *ConcurrentHashMap[K, V]) Remove(key K) (V, bool) {
	h := m.hf(key)
	b := m.table[h%m.capacity]
	b.Lock()
	n := b.remove(h, key, m.ef)
	b.Unlock()
	if n == nil {
		return *new(V), false
	}
	return n.value, true
}

// Size returns the count of entries in the map
func (m *ConcurrentHashMap[K, V]) Size() int {
	var size int64 = 0
	for _, b := range m.table {
		size += b.size
	}
	return int(size)
}

func (b *bucket[K, V]) get(h uint32, key K, ef EqualsFunc[K]) *node[K, V] {
	n := b.node
	for n != nil {
		if n.hash == h && ef(n.key, key) {
			return n
		}
		if b.tree && n.hash > h {
			n = n.left
		} else {
			n = n.right
		}
	}
	return nil
}

func (b *bucket[K, V]) put(h uint32, key K, val V, ef EqualsFunc[K]) {
	if fn := b.get(h, key, ef); fn != nil {
		fn.value = val
		return
	}
	nn := &node[K, V]{
		hash:  h,
		key:   key,
		value: val,
	}
	if b.node == nil {
		b.node = nn
		b.size = 1
		return
	}
	if b.tree {
		if treePut(b.node, nn, ef) {
			b.size++
		}
	} else {
		if listPut(b.node, nn, ef) {
			b.size++
		}
		if b.size >= treeifyThreshold {
			r := treeify(b.node)
			b.node = r
			b.tree = true
		}
	}
}

func (b *bucket[K, V]) remove(h uint32, key K, ef EqualsFunc[K]) (rn *node[K, V]) {
	if b.tree {
		var sn *node[K, V]
		sn, rn = treeRemove(b.node, h, key, ef)
		if rn != nil {
			b.size--
			if b.node == rn {
				b.node = sn
			}
		}
	} else {
		var ok bool
		rn, ok = listRemove(b.node, h, key, ef)
		if ok {
			b.size--
			if rn == nil {
				rn = b.node
				if b.node.right != nil {
					b.node = b.node.right
				} else {
					b.node = nil
				}
			}
		}
	}
	return rn
}

func treeify[K, V any](head *node[K, V]) (root *node[K, V]) {
	nodes := collect(head)
	sort(nodes)
	ri := len(nodes) / 2
	root = nodes[ri]
	split(nodes[:ri], root, true)
	split(nodes[ri+1:], root, false)
	return
}

func split[K, V any](nodes []*node[K, V], root *node[K, V], left bool) {
	l := len(nodes)
	if l == 0 {
		if left {
			root.left = nil
		} else {
			root.right = nil
		}
		return
	}
	ri := len(nodes) / 2
	if left {
		root.left = nodes[ri]
	} else {
		root.right = nodes[ri]
	}
	split(nodes[:ri], nodes[ri], true)
	split(nodes[ri+1:], nodes[ri], false)
}

func collect[K, V any](head *node[K, V]) []*node[K, V] {
	s := size(head)
	ns := make([]*node[K, V], s)
	n := head
	for i := 0; i < s; i++ {
		ns[i] = n
		n = n.right
	}
	return ns
}

func size[K, V any](head *node[V, K]) int {
	n := head
	s := 0
	for n != nil {
		s++
		n = n.right
	}
	return s
}

func sort[K, V any](nodes []*node[K, V]) {
	for i := 0; i < len(nodes)-1; i++ {
		for j := 0; j < len(nodes)-1-i; j++ {
			if nodes[j].hash > nodes[j+1].hash {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}
}

func listRemove[K, V any](n *node[K, V], h uint32, key K, ef EqualsFunc[K]) (*node[K, V], bool) {
	var pn *node[K, V]
	for n != nil {
		if n.hash == h && ef(n.key, key) {
			if pn == nil {
				return nil, true
			}
			pn.right = n.right
			return n, true
		}
		pn = n
		n = n.right
	}
	return nil, false
}

func treeRemove[K, V any](r *node[K, V], h uint32, key K, ef EqualsFunc[K]) (*node[K, V], *node[K, V]) {
	if r == nil {
		return nil, nil
	}
	if r.hash > h {
		var rn *node[K, V]
		r.left, rn = treeRemove(r.left, h, key, ef)
		return r, rn
	} else if r.hash < h || !ef(r.key, key) {
		var rn *node[K, V]
		r.right, rn = treeRemove(r.right, h, key, ef)
		return r, rn
	}
	if r.left == nil {
		return r.right, r
	} else if r.right == nil {
		return r.left, r
	} else {
		spn := r
		sn := r.right
		for sn.left != nil {
			spn = sn
			sn = sn.left
		}
		if spn != r {
			spn.left = sn.right
		} else {
			spn.right = sn.right
		}
		rn := &node[K, V]{
			hash:  r.hash,
			key:   r.key,
			value: r.value,
		}
		r.hash = sn.hash
		r.key = sn.key
		r.value = sn.value
		return r, rn
	}
}

func listPut[K, V any](hn *node[K, V], nn *node[K, V], ef EqualsFunc[K]) bool {
	var pn *node[K, V]
	for hn != nil {
		if hn.hash == nn.hash && ef(hn.key, nn.key) {
			hn.value = nn.value
			return false
		}
		pn = hn
		hn = hn.right
	}
	if pn != nil {
		pn.right = nn
		return true
	}
	return false
}

func treePut[K, V any](rn *node[K, V], nn *node[K, V], ef EqualsFunc[K]) bool {
	var pn *node[K, V]
	for rn != nil {
		if rn.hash == nn.hash && ef(rn.key, nn.key) {
			rn.value = nn.value
			return false
		}
		pn = rn
		if rn.hash > nn.hash {
			rn = rn.left
		} else {
			rn = rn.right
		}
	}
	if pn.hash > nn.hash {
		pn.left = nn
	} else {
		pn.right = nn
	}
	return true
}
