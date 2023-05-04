package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func waitGroup(wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		wg.Add(1) // equals to count (2 times) of goroutines would be run

		go func(id int) {
			defer wg.Done() // alerting the WaitGroup when a goroutine completes

			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d is done\n", id)
		}(i)
	}
}
