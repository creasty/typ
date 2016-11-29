package string_h

import (
	"testing"
)

func Test_Set(t *testing.T) {
	s := NewSet()

	s.Add("a")
	s.Add("b")
	s.Add("b")
	s.Add("c")
	s.Add("c")
	s.Add("c")

	if !s.Include("c") {
		t.Error("error")
	}

	s.Remove("c")
	s.Remove("x")

	if s.Include("c") {
		t.Error("error")
	}

	actual := s.ToArray()
	expect := []string{"a", "b"}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
		return
	}

	for _, e := range expect {
		if !s.Include(e) {
			t.Error("expect to include %v", e)
			break
		}
	}
}
