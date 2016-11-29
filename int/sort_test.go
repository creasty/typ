package int_h

import (
	"testing"
)

func Test_SortInt(t *testing.T) {
	ary := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortInt(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != i+1 {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortInt8(t *testing.T) {
	ary := []int8{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortInt8(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != int8(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortInt16(t *testing.T) {
	ary := []int16{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortInt16(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != int16(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortInt32(t *testing.T) {
	ary := []int32{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortInt32(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != int32(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortInt64(t *testing.T) {
	ary := []int64{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortInt64(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != int64(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortUint(t *testing.T) {
	ary := []uint{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortUint(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != uint(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortUint8(t *testing.T) {
	ary := []uint8{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortUint8(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != uint8(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortUint16(t *testing.T) {
	ary := []uint16{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortUint16(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != uint16(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortUint32(t *testing.T) {
	ary := []uint32{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortUint32(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != uint32(i+1) {
			t.Error("should sort")
			break
		}
	}
}

func Test_SortUint64(t *testing.T) {
	ary := []uint64{1, 9, 2, 8, 3, 7, 4, 6, 5}

	SortUint64(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})

	for i, a := range ary {
		if a != uint64(i+1) {
			t.Error("should sort")
			break
		}
	}
}
