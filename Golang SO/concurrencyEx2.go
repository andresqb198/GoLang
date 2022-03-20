package main

import (
	"fmt"
	"time"
)

func runner1() {
	fmt.Print("\nI am first runner")
}

func runner2() {
	fmt.Print("\nI am second runner")
}

func execute() {
	go runner1()
	go runner2()

}

func main() {

	// Launching both the runners
	execute()
	time.Sleep(time.Second)
}
