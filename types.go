package qingcloud

import (
	"strings"
)

// String 青云的 String 类型
type String struct {
	value string
	enums map[string]bool
	write bool
}

// Set 设置字段属性
func (s *String) Set(t string) {
	if len(s.enums) != 0 {
		if _, ok := s.enums[t]; ok {
			s.value = t
			s.write = true
		}
	} else {
		s.value = t
		s.write = true
	}
}

// Enum 提前设置相关的允许范围
func (s *String) Enum(t ...string) {
	if s.enums == nil {
		s.enums = make(map[string]bool)
		for _, o := range t {
			s.enums[o] = true
		}
	}
	if _, ok := s.enums[s.value]; !ok {
		s.value = ""
		s.write = false
	}
}

func (s *String) String() string {
	return s.value
}

// NumberedInteger X.N
type NumberedInteger struct {
	values map[int64]bool
	enums  map[int64]bool
	write  bool
}

// Add 添加
func (s *NumberedInteger) Add(t ...int64) {
	if s.enums == nil {
		s.enums = make(map[int64]bool)
	}
	if s.values == nil {
		s.values = make(map[int64]bool)
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

// Enum 范围
func (s *NumberedInteger) Enum(e ...int64) {
	if s.enums == nil {
		s.enums = make(map[int64]bool)
	}
	for k := range s.values {
		if _, ok := s.enums[k]; !ok {
			delete(s.values, k)
		}
	}
	if len(s.values) == 0 {
		s.write = false
	}
}

// NumberedString x.N string 类型
type NumberedString struct {
	values map[string]bool
	enums  map[string]bool
	write  bool
}

func (s *NumberedString) String() string {
	var a = []string{}
	for i := range s.values {
		a = append(a, i)
	}
	return strings.Join(a, ",")
}

// Add 添加
func (s *NumberedString) Add(t ...string) {
	if s.enums == nil {
		s.enums = make(map[string]bool)
	}
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

// Enum 范围
func (s *NumberedString) Enum(e ...string) {
	if s.enums == nil {
		s.enums = make(map[string]bool)
		for _, o := range e {
			s.enums[o] = true
		}
	}
	for k := range s.values {
		if _, ok := s.enums[k]; !ok {
			delete(s.values, k)
		}
	}
	if len(s.values) == 0 {
		s.write = false
	}
}

// Integer 请求的 int 类型
type Integer struct {
	value int
	enums map[int]bool
	write bool
}

// Set 设置值
func (s *Integer) Set(t int) {
	if len(s.enums) != 0 {
		if _, ok := s.enums[t]; ok {
			s.value = t
			s.write = true
		}
	} else {
		s.value = t
		s.write = true
	}
}

// Enum 设置范围
func (s *Integer) Enum(e ...int) {
	if s.enums == nil {
		s.enums = make(map[int]bool)
	}
	for _, o := range e {
		s.enums[o] = true
	}
}

// Dict 字典类型
type Dict struct {
	values []map[string]interface{}
}

// Add 添加
func (c *Dict) Add(t map[string]interface{}) {
	if len(c.values) == 0 {
		c.values = make([]map[string]interface{}, 0)
	}
	c.values = append(c.values, t)
}

// Array TODO: 实现这个类型
type Array struct{}
