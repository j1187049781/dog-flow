package util

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete((*s), v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}