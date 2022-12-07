package main

import "fmt"

type qItem[V any] struct {
	value V
	next  *qItem[V]
	prev  *qItem[V]
}
type queue[V any] struct {
	front *qItem[V]
	back  *qItem[V]
}

func mkQueue[V any](value V) *queue[V] {
	out := &queue[V]{}
	firstItem := &qItem[V]{value: value}
	out.front, out.back = firstItem, firstItem
	return out
}

func (q *queue[V]) enqueue(value V) {
	newItem := &qItem[V]{value: value, next: q.back}
	q.back.prev = newItem
	q.back = newItem
}

func (q *queue[V]) dequeue() *qItem[V] {
	out := q.front
	q.front = q.front.next
	return out
}

func (q *queue[V]) print() {
	pos := q.back
	fmt.Println("Back")
	for pos != nil {
		fmt.Println(pos.value)
		pos = pos.next
	}
	fmt.Println("Front")
}
