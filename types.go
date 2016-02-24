package qingcloud

import (
	"strings"
)

type String struct {
	value string
}

func (s *String) Set(t string) {
	s.value = t
}

func (s *String) String() string {
	return s.value
}

type NumberedString struct {
	values map[string]bool
	enums  map[string]bool
}

func (s *NumberedString) String() string {
	var a = []string{}
	for i, _ := range s.values {
		a = append(a, i)
	}
	return strings.Join(a, ",")
}
func (s *NumberedString) Add(t ...string) {
	if s.values == nil {
		s.values = make(map[string]bool)
	}
	for _, o := range t {
		if len(s.enums) != 0 {
			if _, ok := s.enums[o]; ok {
				s.values[o] = true
			}
		} else {
			s.values[o] = true
		}

	}
}

func (s *NumberedString) Enum(e ...string) {
	if s.enums == nil {
		s.enums = make(map[string]bool)
	}
	for _, o := range e {
		s.enums[o] = true
	}
}

type Integer struct {
	value int
	enums map[int]bool
}

func (s *Integer) Set(t int) {

	if len(s.enums) != 0 {
		if _, ok := s.enums[t]; ok {
			s.value = t
		}
	} else {
		s.value = t
	}
}

func (s *Integer) Enum(e ...int) {
	if s.enums == nil {
		s.enums = make(map[int]bool)
	}
	for _, o := range e {
		s.enums[o] = true
	}
}

// TODO 实现这个类型
type NumberedInteger struct {
	values map[string]bool
	enums  map[string]bool
}

// TODO 实现这个类型

type Array struct{}
