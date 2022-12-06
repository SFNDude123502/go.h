package main

import "fmt"

type llNode[V any] struct {
	value V
	next  *llNode[V]
	prev  *llNode[V]
}
type ll[V any] struct {
	head *llNode[V]
	tail *llNode[V]
	len  int
}

func mkLl[V any](value V) *ll[V] {
	headAndTail := &llNode[V]{value: value}
	l := &ll[V]{len: 1}
	l.head, l.tail = headAndTail, headAndTail
	return l
}

func (l *ll[V]) append(value V) {
	newNode := &llNode[V]{value: value, prev: l.tail}
	l.tail.next = newNode
	l.tail = l.tail.next
	l.len++
}

func (l *ll[V]) prepend(value V) {
	newNode := &llNode[V]{value: value, next: l.head}
	l.head = newNode
	l.head.next.prev = l.head
	l.len++
}

func (l *ll[V]) get(index int) *llNode[V] {
	pos := l.head
	for i := 0; pos != nil && i <= index; i++ {
		if i == index {
			return pos
		}
		pos = pos.next
	}
	return nil
}

func (l *ll[V]) remove(index int) {
	node := l.get(index)

	if node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	l.len--
}

func (l *ll[V]) print() {
	pos := l.head
	var out string
	for pos != nil {
		out += fmt.Sprint(pos.value) + " -> "
		pos = pos.next
	}
	fmt.Println(out[:len(out)-3])
}

func (l *ll[V]) iterator() (out []*llNode[V]) {
	pos := l.head
	for pos != nil {
		out = append(out, pos)
		pos = pos.next
	}
	return
}
