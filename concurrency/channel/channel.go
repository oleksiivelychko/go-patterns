package concurrency

import "fmt"

/*
Goroutines have easy communication (with low latency) medium known as channel.
By default, sends and receives block until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.
*/
func channel(ch chan int, s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // send sum to channel
}

/*
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/
func selectChannel(tickCh chan string, quitCh chan int) {
	init := "init"
	for {
		select {
		case tickCh <- init:
			fmt.Println("tick")
		case <-quitCh:
			fmt.Println("quit")
			return // loop exit
		}
	}
}
