package concurrency

import (
	"testing"
)

/*
Non-buffered channels are blocking, meaning that when sending data into a channel, Go waits until the data is read from the channel before execution continues.
Buffered channels, meaning sending data into that channel won’t block until exceed the capacity.
*/
func TestChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int, 2) // make buffered channel

	go channel(ch, s[:len(s)/2])
	go channel(ch, s[len(s)/2:])

	x, y := <-ch, <-ch

	close(ch) // optional

	_, ok := <-ch
	if ok == true {
		t.Fatal("channel has not been closed")
	}

	if x != -5 {
		t.Errorf("value %d is wrong", x)
	}

	if y != 17 {
		t.Errorf("value %d is wrong", x)
	}
}

func TestSelectChannel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	tickCh := make(chan string)
	quitCh := make(chan int)

	go func() {
		t.Log(<-tickCh)
		t.Log(<-tickCh)

		quitCh <- 0
	}()

	selectChannel(tickCh, quitCh)
}
