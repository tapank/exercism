package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

type Set struct {
	v map[string]struct{}
}

func New() Set {
	return Set{make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, s := range l {
		set.v[s] = struct{}{}
	}
	return set
}

func (s Set) String() string {
	var out strings.Builder
	first := true
	out.WriteRune('{')
	for k := range s.v {
		if !first {
			out.WriteString(", ")
		} else {
			first = false
		}
		out.WriteString(fmt.Sprintf("%#v", k))
	}
	out.WriteRune('}')
	return out.String()
}

func (s Set) IsEmpty() bool {
	return len(s.v) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.v[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.v[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k := range s1.v {
		if _, ok := s2.v[k]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1.v {
		if _, ok := s2.v[k]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.v) != len(s2.v) {
		return false
	}
	for k := range s1.v {
		if _, ok := s2.v[k]; !ok {
			return false
		}
	}
	for k := range s2.v {
		if _, ok := s1.v[k]; !ok {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for k := range s1.v {
		if _, ok := s2.v[k]; ok {
			s.Add(k)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for k := range s1.v {
		if _, ok := s2.v[k]; !ok {
			s.Add(k)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for k := range s1.v {
		s.Add(k)
	}
	for k := range s2.v {
		s.Add(k)
	}
	return s
}
