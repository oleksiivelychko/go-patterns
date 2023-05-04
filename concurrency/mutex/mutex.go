package concurrency

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mutex sync.Mutex
	dict  map[string]int
}

// Inc increments the counter for the given key.
func (counter *SafeCounter) Inc(key string) {
	counter.mutex.Lock()
	// Lock so only one goroutine at a time can access the map counter.dict.
	counter.dict[key]++
	counter.mutex.Unlock()
}

// Value returns the current value of the counter for the given key.
func (counter *SafeCounter) Value(key string) int {
	counter.mutex.Lock()
	// Lock so only one goroutine at a time can access the map counter.dict.
	defer counter.mutex.Unlock()
	return counter.dict[key]
}
