package main

import "fmt"

type sItem[V any] struct {
	value V
	below *sItem[V]
}
type stack[V any] struct {
	top *sItem[V]
}

func mkStack[V any](value V) *stack[V] {
	out := &stack[V]{}
	out.top = &sItem[V]{value: value}
	return out
}

func (s *stack[V]) push(value V) {
	newItem := &sItem[V]{value: value, below: s.top}
	s.top = newItem
}
func (s *stack[V]) pop() {
	if s.top != nil {
		s.top = s.top.below
	}
}

func (s *stack[V]) print() {
	pos := s.top
	for pos != nil {
		fmt.Println(pos.value)
		pos = pos.below
	}
	fmt.Println()
}
