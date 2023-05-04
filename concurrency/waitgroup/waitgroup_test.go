package concurrency

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	// is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	waitGroup(&wg)

	// block the execution until the WaitGroup counter goes back to 0, all the workers notified they’re done
	wg.Wait()
}
