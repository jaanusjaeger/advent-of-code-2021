package main

import "fmt"

type stack struct {
	data [100]*pair
	size int
}

func (s *stack) push(p *pair) {
	s.data[s.size] = p
	s.size++
}

func (s *stack) pop() *pair {
	if s.size == 0 {
		return nil
	}
	s.size--
	return s.data[s.size]
}

func (s *stack) peek() *pair {
	if s.size == 0 {
		return nil
	}
	return s.data[s.size-1]
}

func (s *stack) toString() string {
	result := ""
	for i := 0; i < s.size; i++ {
		result += fmt.Sprintf("%v", s.data[i])
	}
	return result
}
