package queue

import (
	"errors"
	"github.com/viodid/ctci-solutions-go/ll"
)

type Queue[T any] struct {
	data *ll.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: &ll.LinkedList[T]{}}
}

func (q *Queue[T]) Add(item T) {
	q.data.AddTail(ll.NewNode(item))
}

func (q *Queue[T]) Remove() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("cannot remove on an empty queue")
	}
	return q.data.RemoveHead().Content, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("cannot peek on an empty queue")
	}
	return q.data.Head.Content, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.data == nil || q.data.Head == nil
}
