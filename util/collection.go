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

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return make(Queue[T], 0)
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) Len() int {
	return len(*q)
}
