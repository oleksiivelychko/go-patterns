package concurrency

import (
	"testing"
	"time"
)

func TestConcurrency_Mutex(t *testing.T) {
	counter := SafeCounter{dict: make(map[string]int)}

	for i := 0; i < 3; i++ {
		go counter.Inc("counter")
	}

	time.Sleep(time.Second)

	if counter.Value("counter") != 3 {
		t.Errorf("value %d is wrong", counter.Value("counter"))
	}
}
