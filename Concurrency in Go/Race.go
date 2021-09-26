package main

// one goroutine is the main
// goroutine that comes by default
import (
	"fmt"
	"runtime"
	"sync"
)

var wgIns sync.WaitGroup

func main() {

	// shared variable
	counter := 0

	// the other 10 goroutines are
	// supposed to come from here
	wgIns.Add(10)
	for i := 0; i < 10; i++ {

		// goroutines are made
		go func() {
			for j := 0; j < 10; j++ {

				// shared variable execution
				counter += 1
				// 100 should be the counter value but
				// it may be equal to 100 or lesser
				// due to race condition
			}
			wgIns.Done()
		}()
	}

	// this value should actually be 11
	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())

	wgIns.Wait()

	// value should be 100
	fmt.Println("Counter value = ", counter)

	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())

	// this value is supposed to be 1
	// but lets see if these values
	// stay consistently same every
	// time we run the code
}
/*This inconsistency happens because of race conditions. Race condition in simple terms can be explained as, you have one candy and two kids run to you claiming that they’re both hungry. They both end up fighting for that one chocolate, they race to get the candy. This is a race condition. Here the solution is: get another candy so that the two enjoy a candy peacefully. Likewise, we can increase the resources allocation in order to ensure race condition does not occur.*/