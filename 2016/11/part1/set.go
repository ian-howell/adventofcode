package main

import "fmt"

type Set[K comparable] map[K]struct{}

func NewSet[K comparable]() Set[K] {
	return Set[K]{}
}

func (s Set[K]) Add(key K) {
	s[key] = struct{}{}
}

func (s Set[K]) Contains(key K) bool {
	_, found := s[key]
	return found
}

func (s Set[K]) Remove(key K) {
	delete(s, key)
}

func (s Set[K]) Len() int {
	return len(s)
}

func (s Set[K]) Members() []K {
	members := []K{}
	for member := range s {
		members = append(members, member)
	}
	return members
}

func (s Set[K]) String() string {
	return fmt.Sprintf("{ %v }", s.Members())
}
