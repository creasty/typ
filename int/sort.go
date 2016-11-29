package int_h

import (
	"sort"
)

/*  int
-----------------------------------------------*/
type sortableInt struct {
	array []int
	less  func(i, j int) bool
}

func SortInt(array []int, less func(i, j int) bool) {
	s := sortableInt{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableInt) Len() int {
	return len(self.array)
}

func (self sortableInt) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableInt) Less(i, j int) bool {
	return self.less(i, j)
}

/*  int8
-----------------------------------------------*/
type sortableInt8 struct {
	array []int8
	less  func(i, j int) bool
}

func SortInt8(array []int8, less func(i, j int) bool) {
	s := sortableInt8{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableInt8) Len() int {
	return len(self.array)
}

func (self sortableInt8) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableInt8) Less(i, j int) bool {
	return self.less(i, j)
}

/*  int16
-----------------------------------------------*/
type sortableInt16 struct {
	array []int16
	less  func(i, j int) bool
}

func SortInt16(array []int16, less func(i, j int) bool) {
	s := sortableInt16{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableInt16) Len() int {
	return len(self.array)
}

func (self sortableInt16) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableInt16) Less(i, j int) bool {
	return self.less(i, j)
}

/*  int32
-----------------------------------------------*/
type sortableInt32 struct {
	array []int32
	less  func(i, j int) bool
}

func SortInt32(array []int32, less func(i, j int) bool) {
	s := sortableInt32{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableInt32) Len() int {
	return len(self.array)
}

func (self sortableInt32) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableInt32) Less(i, j int) bool {
	return self.less(i, j)
}

/*  int64
-----------------------------------------------*/
type sortableInt64 struct {
	array []int64
	less  func(i, j int) bool
}

func SortInt64(array []int64, less func(i, j int) bool) {
	s := sortableInt64{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableInt64) Len() int {
	return len(self.array)
}

func (self sortableInt64) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableInt64) Less(i, j int) bool {
	return self.less(i, j)
}

/*  uint
-----------------------------------------------*/
type sortableUint struct {
	array []uint
	less  func(i, j int) bool
}

func SortUint(array []uint, less func(i, j int) bool) {
	s := sortableUint{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableUint) Len() int {
	return len(self.array)
}

func (self sortableUint) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableUint) Less(i, j int) bool {
	return self.less(i, j)
}

/*  uint8
-----------------------------------------------*/
type sortableUint8 struct {
	array []uint8
	less  func(i, j int) bool
}

func SortUint8(array []uint8, less func(i, j int) bool) {
	s := sortableUint8{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableUint8) Len() int {
	return len(self.array)
}

func (self sortableUint8) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableUint8) Less(i, j int) bool {
	return self.less(i, j)
}

/*  uint16
-----------------------------------------------*/
type sortableUint16 struct {
	array []uint16
	less  func(i, j int) bool
}

func SortUint16(array []uint16, less func(i, j int) bool) {
	s := sortableUint16{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableUint16) Len() int {
	return len(self.array)
}

func (self sortableUint16) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableUint16) Less(i, j int) bool {
	return self.less(i, j)
}

/*  uint32
-----------------------------------------------*/
type sortableUint32 struct {
	array []uint32
	less  func(i, j int) bool
}

func SortUint32(array []uint32, less func(i, j int) bool) {
	s := sortableUint32{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableUint32) Len() int {
	return len(self.array)
}

func (self sortableUint32) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableUint32) Less(i, j int) bool {
	return self.less(i, j)
}

/*  uint64
-----------------------------------------------*/
type sortableUint64 struct {
	array []uint64
	less  func(i, j int) bool
}

func SortUint64(array []uint64, less func(i, j int) bool) {
	s := sortableUint64{
		array: array,
		less:  less,
	}
	sort.Sort(s)
}

func (self sortableUint64) Len() int {
	return len(self.array)
}

func (self sortableUint64) Swap(i, j int) {
	self.array[i], self.array[j] = self.array[j], self.array[i]
}

func (self sortableUint64) Less(i, j int) bool {
	return self.less(i, j)
}
