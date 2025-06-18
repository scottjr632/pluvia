package utils

import "log"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func Invariant(b bool, message string) {
	if !b {
		log.Fatal("Invariant error:", message)
	}
}

func MustCond(b bool, message string) {
	if !b {
		log.Fatal(message)
	}
}
