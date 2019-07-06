package pie

import (
	"sync"
)

// Memory is the in-memory datastore for a pie for development only
type Memory struct {
	total int
	*sync.RWMutex
}

// NewMemory returns a new in-memory datastore
func NewMemory() *Memory {
	return &Memory{
		total:   0,
		RWMutex: &sync.RWMutex{},
	}
}

// Add a slice of pie from the total pie, the slice is split into single dollar sized slices
func (c *Memory) Add(slices int) {
	c.Lock()
	defer c.Unlock()
	c.total += slices

}

// Remove slices of pie from the total pie
func (c *Memory) Remove(slices int) {
	c.Lock()
	defer c.Unlock()
	c.total -= slices
}

// Total slices of pie from the total pie
func (c *Memory) Total() int {
	c.Lock()
	defer c.Unlock()
	return c.total
}
