package main

import "sync"

type Stack[T any] struct {
	data []T
	lock sync.Mutex
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(val T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == 0 {
		var zv T
		return zv, false
	}

	lastIndex := s.lastIndex()
	val := s.data[lastIndex]
	s.data[lastIndex] = *new(T) // clear
	s.data = s.data[:lastIndex]

	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == 0 {
		var zv T
		return zv, false
	}

	return s.data[s.lastIndex()], true
}

func (s *Stack[T]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.data)
}

func (s *Stack[T]) lastIndex() int {
	return len(s.data) - 1
}
