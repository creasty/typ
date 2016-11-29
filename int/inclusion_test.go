package int_h

import (
	"testing"
)

func Test_IncludeInt(t *testing.T) {
	ary := []int{1, 2}

	indexMap := map[int]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeInt(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeInt8(t *testing.T) {
	ary := []int8{1, 2}

	indexMap := map[int8]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeInt8(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeInt16(t *testing.T) {
	ary := []int16{1, 2}

	indexMap := map[int16]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeInt16(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeInt32(t *testing.T) {
	ary := []int32{1, 2}

	indexMap := map[int32]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeInt32(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeInt64(t *testing.T) {
	ary := []int64{1, 2}

	indexMap := map[int64]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeInt64(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeUint(t *testing.T) {
	ary := []uint{1, 2}

	indexMap := map[uint]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeUint(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeUint8(t *testing.T) {
	ary := []uint8{1, 2}

	indexMap := map[uint8]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeUint8(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeUint16(t *testing.T) {
	ary := []uint16{1, 2}

	indexMap := map[uint16]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeUint16(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeUint32(t *testing.T) {
	ary := []uint32{1, 2}

	indexMap := map[uint32]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeUint32(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}

func Test_IncludeUint64(t *testing.T) {
	ary := []uint64{1, 2}

	indexMap := map[uint64]bool{
		1: true,
		2: true,
		9: false,
	}

	for item, found := range indexMap {
		if IncludeUint64(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}
