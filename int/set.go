package int_h

import (
	"sync"
)

/*  int
-----------------------------------------------*/
type IntSet struct {
	set  map[int]bool
	lock sync.Mutex
}

func NewIntSet() *IntSet {
	return &IntSet{
		set:  make(map[int]bool),
		lock: sync.Mutex{},
	}
}

func (self *IntSet) Add(item int) {
	self.set[item] = true
}

func (self *IntSet) AddSafe(item int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *IntSet) Remove(item int) {
	delete(self.set, item)
}

func (self *IntSet) RemoveSafe(item int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *IntSet) ToArray() []int {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]int, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  int8
-----------------------------------------------*/
type Int8Set struct {
	set  map[int8]bool
	lock sync.Mutex
}

func NewInt8Set() *Int8Set {
	return &Int8Set{
		set:  make(map[int8]bool),
		lock: sync.Mutex{},
	}
}

func (self *Int8Set) Add(item int8) {
	self.set[item] = true
}

func (self *Int8Set) AddSafe(item int8) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Int8Set) Remove(item int8) {
	delete(self.set, item)
}

func (self *Int8Set) RemoveSafe(item int8) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Int8Set) ToArray() []int8 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]int8, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  int16
-----------------------------------------------*/
type Int16Set struct {
	set  map[int16]bool
	lock sync.Mutex
}

func NewInt16Set() *Int16Set {
	return &Int16Set{
		set:  make(map[int16]bool),
		lock: sync.Mutex{},
	}
}

func (self *Int16Set) Add(item int16) {
	self.set[item] = true
}

func (self *Int16Set) AddSafe(item int16) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Int16Set) Remove(item int16) {
	delete(self.set, item)
}

func (self *Int16Set) RemoveSafe(item int16) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Int16Set) ToArray() []int16 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]int16, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  int32
-----------------------------------------------*/
type Int32Set struct {
	set  map[int32]bool
	lock sync.Mutex
}

func NewInt32Set() *Int32Set {
	return &Int32Set{
		set:  make(map[int32]bool),
		lock: sync.Mutex{},
	}
}

func (self *Int32Set) Add(item int32) {
	self.set[item] = true
}

func (self *Int32Set) AddSafe(item int32) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Int32Set) Remove(item int32) {
	delete(self.set, item)
}

func (self *Int32Set) RemoveSafe(item int32) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Int32Set) ToArray() []int32 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]int32, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  int64
-----------------------------------------------*/
type Int64Set struct {
	set  map[int64]bool
	lock sync.Mutex
}

func NewInt64Set() *Int64Set {
	return &Int64Set{
		set:  make(map[int64]bool),
		lock: sync.Mutex{},
	}
}

func (self *Int64Set) Add(item int64) {
	self.set[item] = true
}

func (self *Int64Set) AddSafe(item int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Int64Set) Remove(item int64) {
	delete(self.set, item)
}

func (self *Int64Set) RemoveSafe(item int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Int64Set) ToArray() []int64 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]int64, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  uint
-----------------------------------------------*/
type UintSet struct {
	set  map[uint]bool
	lock sync.Mutex
}

func NewUintSet() *UintSet {
	return &UintSet{
		set:  make(map[uint]bool),
		lock: sync.Mutex{},
	}
}

func (self *UintSet) Add(item uint) {
	self.set[item] = true
}

func (self *UintSet) AddSafe(item uint) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *UintSet) Remove(item uint) {
	delete(self.set, item)
}

func (self *UintSet) RemoveSafe(item uint) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *UintSet) ToArray() []uint {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]uint, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  uint8
-----------------------------------------------*/
type Uint8Set struct {
	set  map[uint8]bool
	lock sync.Mutex
}

func NewUint8Set() *Uint8Set {
	return &Uint8Set{
		set:  make(map[uint8]bool),
		lock: sync.Mutex{},
	}
}

func (self *Uint8Set) Add(item uint8) {
	self.set[item] = true
}

func (self *Uint8Set) AddSafe(item uint8) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Uint8Set) Remove(item uint8) {
	delete(self.set, item)
}

func (self *Uint8Set) RemoveSafe(item uint8) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Uint8Set) ToArray() []uint8 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]uint8, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  uint16
-----------------------------------------------*/
type Uint16Set struct {
	set  map[uint16]bool
	lock sync.Mutex
}

func NewUint16Set() *Uint16Set {
	return &Uint16Set{
		set:  make(map[uint16]bool),
		lock: sync.Mutex{},
	}
}

func (self *Uint16Set) Add(item uint16) {
	self.set[item] = true
}

func (self *Uint16Set) AddSafe(item uint16) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Uint16Set) Remove(item uint16) {
	delete(self.set, item)
}

func (self *Uint16Set) RemoveSafe(item uint16) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Uint16Set) ToArray() []uint16 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]uint16, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  uint32
-----------------------------------------------*/
type Uint32Set struct {
	set  map[uint32]bool
	lock sync.Mutex
}

func NewUint32Set() *Uint32Set {
	return &Uint32Set{
		set:  make(map[uint32]bool),
		lock: sync.Mutex{},
	}
}

func (self *Uint32Set) Add(item uint32) {
	self.set[item] = true
}

func (self *Uint32Set) AddSafe(item uint32) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Uint32Set) Remove(item uint32) {
	delete(self.set, item)
}

func (self *Uint32Set) RemoveSafe(item uint32) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Uint32Set) ToArray() []uint32 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]uint32, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}

/*  uint64
-----------------------------------------------*/
type Uint64Set struct {
	set  map[uint64]bool
	lock sync.Mutex
}

func NewUint64Set() *Uint64Set {
	return &Uint64Set{
		set:  make(map[uint64]bool),
		lock: sync.Mutex{},
	}
}

func (self *Uint64Set) Add(item uint64) {
	self.set[item] = true
}

func (self *Uint64Set) AddSafe(item uint64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Uint64Set) Remove(item uint64) {
	delete(self.set, item)
}

func (self *Uint64Set) RemoveSafe(item uint64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Uint64Set) ToArray() []uint64 {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]uint64, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}
