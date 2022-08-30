package hashmap

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type customType struct {
	i int
	s string
}

func (c customType) Hash() uint32 {
	return intHashFunc(c.i) + stringHashFunc(c.s)
}

func (c customType) Equals(a any) bool {
	ct, ok := a.(customType)
	if !ok {
		return false
	}
	return c.i == ct.i && c.s == ct.s
}

func TestNew(t *testing.T) {
	hf := func(key int) uint32 {
		return uint32(key)
	}
	ef := func(k1, k2 int) bool {
		return k1 == k2
	}
	_ = NewHashMap[int, int](hf, ef)
}

func TestNewWithCap(t *testing.T) {
	hf := func(key int) uint32 {
		return uint32(key)
	}
	ef := func(k1, k2 int) bool {
		return k1 == k2
	}

	_, err := NewWithCap[int, int](32, hf, ef)
	if err != nil {
		t.FailNow()
	}

	_, err = NewWithCap[int, int](0, hf, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewWithCap[int, int](-1, hf, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewWithCap[int, int](32, nil, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewWithCap[int, int](32, hf, nil)
	if err == nil {
		t.FailNow()
	}
}

func TestNewMap(t *testing.T) {
	m := NewMap[customType, int]()
	m.Put(customType{9, "a"}, 9)
}

func TestNewMapWithCap(t *testing.T) {
	_, err := NewMapWithCap[Hasher, int](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewMapWithCap[Hasher, int](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewMapWithCap[Hasher, int](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestNewComparableMap(t *testing.T) {
	m := NewComparableMap[string, int]()
	m.Put("a", 9)
}

func TestNewComparableMapWithCap(t *testing.T) {
	_, err := NewComparableMapWithCap[string, int](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewComparableMapWithCap[string, int](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewComparableMapWithCap[string, int](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestNewStringMap(t *testing.T) {
	m := NewStringMap[int]()
	m.Put("a", 9)
}

func TestNewStringMapWithCap(t *testing.T) {
	_, err := NewStringMapWithCap[int](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewStringMapWithCap[int](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewStringMapWithCap[int](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestNewIntMap(t *testing.T) {
	m := NewIntMap[string]()
	m.Put(9, "a")
}

func TestNewIntMapWithCap(t *testing.T) {
	_, err := NewIntMapWithCap[int](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewIntMapWithCap[int](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewIntMapWithCap[int](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestConcurrentHashMap_Put(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 5_000; i < 15_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
}

func TestConcurrentHashMap_Put_Verify(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 1_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 500; i < 1_500; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	verifyMap(t, &m)
}

func TestConcurrentHashMap_ConcurrentlyPut(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 10_000)
	go putRange(&m, &wg, 5_000, 15_000)
	go putRange(&m, &wg, 10_000, 20_000)
	wg.Wait()
}

func TestConcurrentHashMap_ConcurrentlyPut_Verify(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 1_000)
	go putRange(&m, &wg, 500, 1_500)
	go putRange(&m, &wg, 1_000, 2_000)
	wg.Wait()
	verifyMap(t, &m)
}

func TestConcurrentHashMap_Put_Get(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		m.Put(k, i)
		if v, ok := m.Get(k); !ok || v != i {
			t.Logf("key: %s, value: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
	for i := 10_000; i < 15_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Get(k); ok {
			t.Logf("key: %s, value: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_Get(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 10_000)
	go putRange(&m, &wg, 5_000, 15_000)
	go putRange(&m, &wg, 10_000, 20_000)
	wg.Wait()
	for i := 0; i < 20_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Get(k); !ok || v != i {
			t.Logf("key: %s, value: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_GetOrDefault(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		m.Put(k, i)
		if v := m.GetOrDefault(k, -1); v != i {
			t.Logf("key: %s, value: %v", k, v)
			t.FailNow()
		}
	}
	for i := 10_000; i < 15_000; i++ {
		k := strconv.Itoa(i)
		if v := m.GetOrDefault(k, -1); v != -1 {
			t.Logf("key: %s, value: %v", k, v)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_GetOrDefault(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 10_000)
	go putRange(&m, &wg, 5_000, 15_000)
	go putRange(&m, &wg, 10_000, 20_000)
	wg.Wait()
	for i := 0; i < 20_000; i++ {
		k := strconv.Itoa(i)
		if v := m.GetOrDefault(k, -1); v != i {
			t.Logf("key: %s, value: %v", k, v)
			t.FailNow()
		}
	}
	for i := 20_000; i < 25_000; i++ {
		k := strconv.Itoa(i)
		if v := m.GetOrDefault(k, -1); v != -1 {
			t.Logf("key: %s, value: %v", k, v)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_Contains(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		m.Put(k, i)
		if ok := m.Contains(k); !ok {
			t.Logf("key: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
	for i := 10_000; i < 15_000; i++ {
		k := strconv.Itoa(i)
		if ok := m.Contains(k); ok {
			t.Logf("key: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_Contains(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 10_000)
	go putRange(&m, &wg, 5_000, 15_000)
	go putRange(&m, &wg, 10_000, 20_000)
	wg.Wait()
	for i := 0; i < 20_000; i++ {
		k := strconv.Itoa(i)
		if ok := m.Contains(k); !ok {
			t.Logf("key: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
	for i := 20_000; i < 25_000; i++ {
		k := strconv.Itoa(i)
		if ok := m.Contains(k); ok {
			t.Logf("key: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_Remove(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); !ok && v != i {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); ok {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_Remove_Verify(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 1_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < 1_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); !ok && v != i {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
		if i+1%200 == 0 {
			verifyMap(t, &m)
		}
	}
	for i := 0; i < 1_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); ok {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_Remove(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 10_000)
	go putRange(&m, &wg, 5_000, 15_000)
	go putRange(&m, &wg, 10_000, 20_000)
	wg.Wait()
	for i := 0; i < 20_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); !ok || v != i {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
	for i := 0; i < 20_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); ok {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_Remove_Verify(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putRange := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putRange(&m, &wg, 0, 1_000)
	go putRange(&m, &wg, 500, 1_500)
	go putRange(&m, &wg, 1_000, 2_000)
	wg.Wait()
	for i := 0; i < 2_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); !ok || v != i {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
		if i+1%500 == 0 {
			verifyMap(t, &m)
		}
	}
	for i := 0; i < 2_000; i++ {
		k := strconv.Itoa(i)
		if v, ok := m.Remove(k); ok {
			t.Logf("K: %s, V: %v, ok: %v", k, v, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_Remove_Contains(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		k := strconv.Itoa(i)
		m.Put(k, i)
		if val, ok := m.Remove(k); !ok || val != i {
			t.FailNow()
		}
		if ok := m.Contains(k); ok {
			t.Logf("K: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
	for i := 10_000; i < 15_000; i++ {
		k := strconv.Itoa(i)
		if _, ok := m.Remove(k); ok {
			t.Logf("K: %s, ok: %v", k, ok)
			t.FailNow()
		}
	}
}

func TestConcurrentHashMap_Put_Size(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	if s := m.Size(); s != 10_000 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
	for i := 5_000; i < 15_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	if s := m.Size(); s != 15_000 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
}

func TestConcurrentHashMap_ConcurrentlyPut_Size(t *testing.T) {
	m := NewStringMap[int]()
	var wg sync.WaitGroup
	putItr := func(m *ConcurrentHashMap[string, int], wg *sync.WaitGroup, from, to int) {
		defer wg.Done()
		for i := from; i < to; i++ {
			m.Put(strconv.Itoa(i), i)
		}
	}
	wg.Add(3)
	go putItr(&m, &wg, 0, 10_000)
	go putItr(&m, &wg, 5_000, 15_000)
	go putItr(&m, &wg, 10_000, 20_000)
	wg.Wait()
	if s := m.Size(); s != 20_000 {
		t.Errorf("size: %d", s)
		t.FailNow()
	}
}

func TestConcurrentHashMap_Put_Remove_Size(t *testing.T) {
	m := NewStringMap[int]()
	for i := 0; i < 10_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	if s := m.Size(); s != 10_000 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
	for i := 5_000; i < 15_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	if s := m.Size(); s != 15_000 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
	for i := 10_000; i < 20_000; i++ {
		m.Remove(strconv.Itoa(i))
	}
	if s := m.Size(); s != 10_000 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
	for i := 0; i < 10_000; i++ {
		m.Remove(strconv.Itoa(i))
	}
	if s := m.Size(); s != 0 {
		t.Logf("size: %d", s)
		t.FailNow()
	}
}

func TestTreeify(t *testing.T) {
	n := &node[string, int]{
		hash: rand.Uint32(),
	}
	head := n
	for i := 0; i < 9; i++ {
		n.right = &node[string, int]{
			hash: rand.Uint32(),
		}
		n = n.right
	}
	root := treeify(head)
	if !treeRootNodeVerified(root) {
		t.FailNow()
	}
}

// Utility functions

func verifyMap[K, V any](t *testing.T, m *ConcurrentHashMap[K, V]) {
	if !mapVerified(m) {
		t.Log("Unverified map")
		t.FailNow()
	}
}

func mapVerified[K, V any](m *ConcurrentHashMap[K, V]) bool {
	for i := 0; i < int(m.capacity); i++ {
		if b := m.table[i]; b.tree {
			if !treeRootNodeVerified(b.node) {
				return false
			}
		}
	}
	return true
}

func treeRootNodeVerified[K, V any](r *node[K, V]) bool {
	if r == nil {
		return true
	}
	return treeLeafNodeVerified(r) && treeRootNodeVerified(r.left) && treeRootNodeVerified(r.right)
}

func treeLeafNodeVerified[K, V any](l *node[K, V]) bool {
	if l == nil {
		return true
	}
	if l.right != nil && l.left != nil {
		return l.hash < l.right.hash && l.hash > l.left.hash && l.right.hash > l.left.hash
	} else if l.right != nil {
		return l.hash < l.right.hash
	} else if l.left != nil {
		return l.hash > l.left.hash
	} else {
		return true
	}
}

// Benchmark tests

func BenchmarkConcurrentHashMap_Put(b *testing.B) {
	m := NewStringMap[int]()
	for i := 0; i < b.N; i++ {
		m.Put(strconv.Itoa(i), i)
	}
}

func BenchmarkConcurrentHashMap_Get(b *testing.B) {
	m := NewStringMap[int]()
	for i := 0; i < 100_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < b.N; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func BenchmarkConcurrentHashMap_GetOrDefault(b *testing.B) {
	m := NewStringMap[int]()
	for i := 0; i < 100_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < b.N; i++ {
		m.GetOrDefault(strconv.Itoa(i), -1)
	}
}

func BenchmarkConcurrentHashMap_Contains(b *testing.B) {
	m := NewStringMap[int]()
	for i := 0; i < 100_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < b.N; i++ {
		m.Contains(strconv.Itoa(i))
	}
}

func BenchmarkConcurrentHashMap_Remove(b *testing.B) {
	m := NewStringMap[int]()
	for i := 0; i < 100_000; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	for i := 0; i < b.N; i++ {
		m.Remove(strconv.Itoa(i))
	}
}

func BenchmarkTreeify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := &node[string, int]{
			hash: rand.Uint32(),
		}
		head := n
		for i := 0; i < defaultCapacity; i++ {
			n.right = &node[string, int]{
				hash: rand.Uint32(),
			}
			n = n.right
		}
		treeify(head)
	}
}
