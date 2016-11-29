package math_h

import (
	"math"
)

/*  int
-----------------------------------------------*/
func MinInt(values ...int) int {
	min := math.MaxInt32

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxInt(values ...int) int {
	max := math.MinInt32

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  int8
-----------------------------------------------*/
func MinInt8(values ...int8) int8 {
	min := int8(math.MaxInt8)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxInt8(values ...int8) int8 {
	max := int8(math.MinInt8)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  int16
-----------------------------------------------*/
func MinInt16(values ...int16) int16 {
	min := int16(math.MaxInt16)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxInt16(values ...int16) int16 {
	max := int16(math.MinInt16)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  int32
-----------------------------------------------*/
func MinInt32(values ...int32) int32 {
	min := int32(math.MaxInt32)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxInt32(values ...int32) int32 {
	max := int32(math.MinInt32)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  int64
-----------------------------------------------*/
func MinInt64(values ...int64) int64 {
	min := int64(math.MaxInt64)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxInt64(values ...int64) int64 {
	max := int64(math.MinInt64)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  uint
-----------------------------------------------*/
func MinUint(values ...uint) uint {
	min := uint(math.MaxUint32)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxUint(values ...uint) uint {
	max := uint(0)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  uint8
-----------------------------------------------*/
func MinUint8(values ...uint8) uint8 {
	min := uint8(math.MaxUint8)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxUint8(values ...uint8) uint8 {
	max := uint8(0)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  uint16
-----------------------------------------------*/
func MinUint16(values ...uint16) uint16 {
	min := uint16(math.MaxUint16)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxUint16(values ...uint16) uint16 {
	max := uint16(0)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  uint32
-----------------------------------------------*/
func MinUint32(values ...uint32) uint32 {
	min := uint32(math.MaxUint32)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxUint32(values ...uint32) uint32 {
	max := uint32(0)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

/*  uint64
-----------------------------------------------*/
func MinUint64(values ...uint64) uint64 {
	min := uint64(math.MaxUint64)

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func MaxUint64(values ...uint64) uint64 {
	max := uint64(0)

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}
