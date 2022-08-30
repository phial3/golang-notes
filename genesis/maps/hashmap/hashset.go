package hashmap

type ConcurrentHashSet[T any] struct {
	hashmap ConcurrentHashMap[T, any]
}

func NewSet[T any](hf HashFunc[T], ef EqualsFunc[T]) ConcurrentHashSet[T] {
	return ConcurrentHashSet[T]{
		NewHashMap[T, any](hf, ef),
	}
}

func NewSetWithCap[T any](capacity int, hf HashFunc[T], ef EqualsFunc[T]) (ConcurrentHashSet[T], error) {
	var hashset ConcurrentHashSet[T]
	hashmap, err := NewWithCap[T, any](capacity, hf, ef)
	if err != nil {
		return hashset, err
	}
	hashset.hashmap = hashmap
	return hashset, err
}

func NewHasherSet[T Hasher]() ConcurrentHashSet[T] {
	return ConcurrentHashSet[T]{
		NewMap[T, any](),
	}
}

func NewHasherSetWithCap[T Hasher](capacity int) (ConcurrentHashSet[T], error) {
	var hashset ConcurrentHashSet[T]
	hashmap, err := NewMapWithCap[T, any](capacity)
	if err != nil {
		return hashset, err
	}
	hashset.hashmap = hashmap
	return hashset, err
}

func NewComparableSet[T comparable]() ConcurrentHashSet[T] {
	return ConcurrentHashSet[T]{
		NewComparableMap[T, any](),
	}
}

func NewComparableSetWithCap[T comparable](capacity int) (ConcurrentHashSet[T], error) {
	var hashset ConcurrentHashSet[T]
	hashmap, err := NewComparableMapWithCap[T, any](capacity)
	if err != nil {
		return hashset, err
	}
	hashset.hashmap = hashmap
	return hashset, err
}

func (s *ConcurrentHashSet[T]) Put(t T) {
	s.hashmap.Put(t, nil)
}

func (s *ConcurrentHashSet[T]) Contains(t T) bool {
	return s.hashmap.Contains(t)
}

func (s *ConcurrentHashSet[T]) Remove(t T) bool {
	_, ok := s.hashmap.Remove(t)
	return ok
}

func (s *ConcurrentHashSet[T]) size() int {
	return s.hashmap.Size()
}
