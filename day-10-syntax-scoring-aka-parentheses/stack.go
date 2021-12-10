package main

type stack struct {
	data [200]rune
	size int
}

func (s *stack) push(r rune) {
	s.data[s.size] = r
	s.size++
}

func (s *stack) pop() rune {
	if s.size == 0 {
		return 'X'
	}
	s.size--
	return s.data[s.size]
}

func (s *stack) toString() string {
	result := ""
	for i := 0; i < s.size; i++ {
		result += string(s.data[i])
	}
	return result
}
