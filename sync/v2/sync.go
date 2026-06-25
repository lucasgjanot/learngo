package sync

import (
	"sync/atomic"
)

type Counter struct {
	value atomic.Int32
}

func (c *Counter) Inc() {
	c.value.Add(1)
}

func (c *Counter) Value() int {
	return int(c.value.Load())
}

func NewCounter() *Counter {
	return &Counter{}
}