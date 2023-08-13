package core

import "sync"

// Container ...
type Container struct {
	hints []string
	m     sync.Mutex
}

// AddHint to Container.
func (c *Container) AddHint(hint string) {
	if c == nil {
		return
	}

	c.m.Lock()
	defer c.m.Unlock()

	c.hints = append(c.hints, hint)
}

// SetHints in Container.
func (c *Container) SetHints(hints []string) {
	if c == nil {
		return
	}

	c.m.Lock()
	defer c.m.Unlock()

	c.hints = hints
}

// GetHints from Container.
func (c *Container) GetHints() []string {
	if c == nil {
		return nil
	}

	c.m.Lock()
	defer c.m.Unlock()

	return c.hints
}
