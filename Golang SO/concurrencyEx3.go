package main

import (
	"fmt"
	"sync"
)

func runner(wg *sync.WaitGroup, runner int) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nNumber ", runner)

}

/*
func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nNumber 2")
}
*/
func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner(wg, 1)
	go runner(wg, 2)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}
