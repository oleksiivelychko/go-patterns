package concurrency

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread managed by the Go runtime.
A goroutine an abstraction over threads, a single OS thread can run many goroutines.
A goroutine takes about 2kb of stack space to initialize.
Goroutines run in the same address space, so access to shared memory must be synchronized.
Goroutine are not hardware dependent and does not have ID because go does not have Thread Local Storage.
*/
func goroutine(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
