package concurrency

import (
	"runtime"
	"testing"
)

func TestGoroutine(t *testing.T) {
	// determines how many threads that program will use simultaneously
	// swapping between threads is a relatively slow operation
	runtime.GOMAXPROCS(2)

	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	go goroutine("async")
	goroutine("sync")
}
