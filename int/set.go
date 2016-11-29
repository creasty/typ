package int_h

import (
	"sync"
)

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
	self.lock.Lock()
	defer self.lock.Unlock()

	self.set[item] = true
}

func (self *IntSet) Remove(item int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	delete(self.set, item)
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
