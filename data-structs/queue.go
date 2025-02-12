package main

import "sync"

type Queue[T any] struct {
	data []T
	lock sync.Mutex
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) == 0 {
		var zv T
		return zv, false
	}

	val := q.data[0]
	q.data[0] = *new(T) // clear
	q.data = q.data[1:]

	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) == 0 {
		var zv T
		return zv, false
	}

	return q.data[0], true
}

func (q *Queue[T]) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.data)
}
