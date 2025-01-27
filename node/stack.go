package node

import "cmp"

type stack[T cmp.Ordered] struct {
	data []*node[T]
	l    int
}

func NewStack[T cmp.Ordered](capacity int) *stack[T] {
	return &stack[T]{
		data: make([]*node[T], 0, capacity),
	}
}

func (s stack[T]) Len() int {
	return s.l
}

func (s *stack[T]) Push(n *node[T]) {
	s.data = append(s.data, n)
	s.l += 1
}

func (s *stack[T]) Pop() *node[T] {
	if s.l == 0 {
		return nil
	}
	ret := s.data[s.l-1]
	s.data = s.data[:s.l-1]
	s.l -= 1

	return ret
}

func (s *stack[T]) Peep() *node[T] {
	if s.l == 0 {
		return nil
	}
	return s.data[s.l-1]
}
