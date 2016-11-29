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

func Test_Int8Set(t *testing.T) {
	s := NewInt8Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []int8{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Int16Set(t *testing.T) {
	s := NewInt16Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []int16{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Int32Set(t *testing.T) {
	s := NewInt32Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []int32{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Int64Set(t *testing.T) {
	s := NewInt64Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []int64{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_UintSet(t *testing.T) {
	s := NewUintSet()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []uint{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Uint8Set(t *testing.T) {
	s := NewUint8Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []uint8{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Uint16Set(t *testing.T) {
	s := NewUint16Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []uint16{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Uint32Set(t *testing.T) {
	s := NewUint32Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []uint32{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}

func Test_Uint64Set(t *testing.T) {
	s := NewUint64Set()

	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Add(3)

	s.Remove(3)
	s.Remove(9)

	actual := s.ToArray()
	expect := []uint64{1, 2}

	if len(actual) != len(expect) {
		t.Error("expect %v to equal to %v", actual, expect)
	}
}
