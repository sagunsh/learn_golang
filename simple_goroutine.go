package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)     // create a channel to pass ints
	cpus := runtime.NumCPU() // number of available cpus

	fmt.Println("number of cpus =", cpus)
	for i := 0; i < cpus; i++ {
		go init_routine(i, ch) // start goroutine
	}

	for i := 0; i < cpus; i++ {
		id := <-ch // receive a value from a channel
		fmt.Println("thread", id, "finished")
	}
	// all goroutines are finished at this point
}

func init_routine(id int, c chan int) {
	fmt.Println("thread", id, "started")
	time.Sleep(4 * time.Second) // sleep
	c <- id                     // send a value back to main
}
