package int_h

import (
	"testing"
)

func Test_IntSet(t *testing.T) {
	s := NewIntSet()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []int{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}
