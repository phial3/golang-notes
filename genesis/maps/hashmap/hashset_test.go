package hashmap

import "testing"

func TestNewSet(t *testing.T) {
	hf := func(key int) uint32 {
		return uint32(key)
	}
	ef := func(k1, k2 int) bool {
		return k1 == k2
	}
	_ = NewSet[int](hf, ef)
}

func TestNewSetWithCap(t *testing.T) {
	hf := func(key int) uint32 {
		return uint32(key)
	}
	ef := func(k1, k2 int) bool {
		return k1 == k2
	}

	_, err := NewSetWithCap[int](32, hf, ef)
	if err != nil {
		t.FailNow()
	}

	_, err = NewSetWithCap[int](0, hf, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewSetWithCap[int](-1, hf, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewSetWithCap[int](32, nil, ef)
	if err == nil {
		t.FailNow()
	}

	_, err = NewSetWithCap[int](32, hf, nil)
	if err == nil {
		t.FailNow()
	}
}

func TestNewHasherSet(t *testing.T) {
	s := NewHasherSet[customType]()
	s.Put(customType{9, "a"})
}

func TestNewHasherSetWithCap(t *testing.T) {
	_, err := NewHasherSetWithCap[Hasher](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewHasherSetWithCap[Hasher](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewHasherSetWithCap[Hasher](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestNewComparableSet(t *testing.T) {
	s := NewComparableSet[string]()
	s.Put("a")
}

func TestNewComparableSetWithCap(t *testing.T) {
	_, err := NewComparableSetWithCap[string](32)
	if err != nil {
		t.FailNow()
	}

	_, err = NewComparableSetWithCap[string](0)
	if err == nil {
		t.FailNow()
	}

	_, err = NewComparableSetWithCap[string](-1)
	if err == nil {
		t.FailNow()
	}
}

func TestConcurrentHashSet_Put(t *testing.T) {
	s := NewComparableSet[int]()
	for i := 0; i < 10_000; i++ {
		s.Put(i)
	}
}

func TestConcurrentHashSet_Remove(t *testing.T) {
	s := NewComparableSet[int]()
	for i := 0; i < 10_000; i++ {
		s.Put(i)
	}
	for i := 5_000; i < 15_000; i++ {
		if i < 10_000 {
			if ok := s.Remove(i); !ok {
				t.Logf("data: %v, ok: %v", i, ok)
				t.FailNow()
			}
		} else {
			if ok := s.Remove(i); ok {
				t.Logf("data: %v, ok: %v", i, ok)
				t.FailNow()
			}
		}
	}
}

func TestConcurrentHashSet_Contains(t *testing.T) {
	s := NewComparableSet[int]()
	for i := 0; i < 10_000; i++ {
		s.Put(i)
	}
	for i := 5_000; i < 15_000; i++ {
		if i < 10_000 {
			if ok := s.Contains(i); !ok {
				t.Logf("data: %v, ok: %v", i, ok)
				t.FailNow()
			}
		} else {
			if ok := s.Contains(i); ok {
				t.Logf("data: %v, ok: %v", i, ok)
				t.FailNow()
			}
		}
	}
}

func TestConcurrentHashMap_Size(t *testing.T) {
	s := NewComparableSet[int]()
	for i := 0; i < 10_000; i++ {
		s.Put(i)
	}
	for i := 5_000; i < 15_000; i++ {
		s.Remove(i)
	}
	if size := s.size(); size != 5000 {
		t.Logf("size: %d", size)
		t.FailNow()
	}
}
