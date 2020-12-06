package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	n               int
	leftCS, rightCS *Chopstick
}

// type Host struct {
//     hostc chan int
// }

var hostc = make(chan int, 2)
var wg sync.WaitGroup
var chan1 = make(chan int)
var chan2 = make(chan int)
var chan3 = make(chan int)
var chan4 = make(chan int)
var chan5 = make(chan int)
var sli = [](chan int){chan1, chan2, chan3, chan4, chan5}

func (p Philosopher) AskPermission() {
	hostc <- 1
	sli[p.n-1] <- 1
	wg.Done()
}

func (p Philosopher) Eat() {
	<-sli[p.n-1]
	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Printf("startng to eat %d\n", p.n)
	fmt.Printf("finishing eating %d\n", p.n)

	p.rightCS.Unlock()
	p.leftCS.Unlock()
	<-hostc
	wg.Done()
}

func main() {

	// Initialization

	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	// Eating

	wg.Add(30)
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			go philosophers[i].AskPermission()
			go philosophers[i].Eat()
		}
	}
	wg.Wait()

}
