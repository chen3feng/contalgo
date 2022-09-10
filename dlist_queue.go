package stl4go

import (
	"fmt"
)

// DListQueue is a FIFO container
type DListQueue[T any] struct {
	list DList[T]
}

// NewDListQueue create a new Queue object.
func NewDListQueue[T any]() *DListQueue[T] {
	q := DListQueue[T]{}
	return &q
}

// Len implements the Container interface.
func (q *DListQueue[T]) Len() int {
	return q.list.Len()
}

// IsEmpty implements the Container interface.
func (q *DListQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear implements the Container interface.
func (q *DListQueue[T]) Clear() {
	q.list.Clear()
}

// Len implements the fmt.Stringer interface.
func (q *DListQueue[T]) String() string {
	return fmt.Sprintf("Queue[%v]", nameOfType[T]())
}

// Front returns the first element in the container.
func (q *DListQueue[T]) Front() T {
	return q.list.Front()
}

// Back returns the last element in the container.
func (q *DListQueue[T]) Back() T {
	return q.list.Back()
}

// PushFront pushed an element to the front of the queue.
func (q *DListQueue[T]) PushFront(val T) {
	q.list.PushFront(val)
}

// PushBack pushed an element to the back of the queue.
func (q *DListQueue[T]) PushBack(val T) {
	q.list.PushBack(val)
}

// PopFront popups an element from the front of the queue.
func (q *DListQueue[T]) PopFront() T {
	return q.list.PopFront()
}

// PopBack popups an element from the back of the queue.
func (q *DListQueue[T]) PopBack() T {
	return q.list.PopBack()
}

// TryPopFront tries popuping an element from the front of the queue.
func (q *DListQueue[T]) TryPopFront() (T, bool) {
	return q.list.TryPopFront()
}

// TryPopBack tries popuping an element from the back of the queue.
func (q *DListQueue[T]) TryPopBack() (T, bool) {
	return q.list.TryPopBack()
}
