package main

import (
	"fmt"
	"sync"
)

func AddOne(wg *sync.WaitGroup, x *int) {
	*x = *x + 1
	wg.Done()
}

func PrintX(wg *sync.WaitGroup, x int) {
	fmt.Println(x)
	wg.Done()
}

func main() {
	var x int = 1
	var wg sync.WaitGroup
	wg.Add(2)
	go AddOne(&wg, &x)
	go PrintX(&wg, x)
	wg.Wait()
}

/*
The race condition here is that the function PrintX can print out 1 or 2 depending on whether AddOne is executed before or after the printing. Since this is a non-deterministic output, it is a race condition.
*/
