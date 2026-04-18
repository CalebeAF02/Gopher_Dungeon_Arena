package goguide

import "sync"

// Counter demonstra como proteger acesso concorrente a um valor com mutex.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment aumenta o contador de forma segura.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value retorna o valor atual do contador com segurança.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
