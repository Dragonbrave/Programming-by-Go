package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)
var checkl, checkr int
var amount int = 5

// a single chopstick can be used as a mutex object
type ChopS struct{ sync.Mutex }

// philosopher has an ID, and 2 assigned chopsticks
type Philo struct {
	id      int
	leftCS  *ChopS
	rightCS *ChopS
}

// philosopher's eat function.
// a philosopher is supposed to eat only 3 times.
// Moreover, the host is supposed to control how many philosophers can eat the the same time (2 here)

func (p Philo) eat() {

	// eat for 3 times
	for i := 0; i < 3; i++ {

		//get permission from the host
		// if host chan is full, wait here
		host <- true

		// get random hold of the mutexes for chopsticks
		var left bool
		if checkr==amount-1 || (rand.Intn(2) == 0 && checkl != amount-1) {
			checkl++
			p.leftCS.Lock()
			p.rightCS.Lock()
			left = true
		} else {
			checkr++
			p.rightCS.Lock()
			p.leftCS.Lock()
			left = false
		}

		fmt.Println("Starting to eat: ", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Println("finishing eating: ", p.id+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		// let host know that the eating is done.
		// free host chan
		<-host
		if left == true {
			checkl--
		} else {
			checkr--
		}
	}
	wg.Done()
}

func main() {

	//create the chopsticks
	ChopSticks := make([]*ChopS, 5)
	for i := 0; i < amount; i++ {
		ChopSticks[i] = new(ChopS)
	}

	// create the philosophers
	Philosophers := make([]*Philo, 5)
	for i := 0; i < amount; i++ {
		Philosophers[i] = &Philo{i, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	// let the philosophers eat now
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}

	wg.Wait()

}
