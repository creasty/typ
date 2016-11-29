package string_h

import (
	"sync"
)

type Set struct {
	set  map[string]bool
	lock sync.Mutex
}

func NewSet() *Set {
	return &Set{
		set:  make(map[string]bool),
		lock: sync.Mutex{},
	}
}

func (self *Set) Add(item string) {
	self.set[item] = true
}

func (self *Set) AddSafe(item string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Add(item)
}

func (self *Set) Remove(item string) {
	delete(self.set, item)
}

func (self *Set) RemoveSafe(item string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.Remove(item)
}

func (self *Set) Include(item string) bool {
	self.lock.Lock()
	defer self.lock.Unlock()

	_, ok := self.set[item]
	return ok
}

func (self *Set) ToArray() []string {
	self.lock.Lock()
	defer self.lock.Unlock()

	a := make([]string, len(self.set))

	i := 0
	for item := range self.set {
		a[i] = item
		i++
	}

	return a
}
