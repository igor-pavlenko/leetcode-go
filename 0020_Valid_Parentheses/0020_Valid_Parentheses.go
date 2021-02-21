package leetcode_0020_Valid_Parentheses

import (
	"errors"
)

func isValid(s string) bool {
	ob := []byte("({[")
	cb := []byte(")}]")
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == ob[0] || s[i] == ob[1] || s[i] == ob[2] {
			stack.Push(s[i])
		} else {
			prev, err := stack.Pop()
			if err != nil {
				return false
			}
			if (s[i] == cb[0] && prev != ob[0]) ||
				(s[i] == cb[1] && prev != ob[1]) ||
				(s[i] == cb[2] && prev != ob[2]) {
				return false
			}
		}
	}
	if _, err := stack.Pop(); err == nil {
		return false
	}
	return true
}

type stack struct {
	s    []byte
}

func NewStack() *stack {
	return &stack{make([]byte, 0)}
}

func (s *stack) Push(v byte) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (byte, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty Stack")
	}
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
