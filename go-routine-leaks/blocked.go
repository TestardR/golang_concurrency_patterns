package goroutineleaks

import (
	"crypto/rand"
	"fmt"
)

// The deferred fmt.Println statement never gets run.
// After the third iteration of our loop, our goroutine blocks trying to send the next random integer to a channel that is no longer being read from. We have no way of telling
// the producer it can stop.

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
