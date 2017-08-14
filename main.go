package main

import (
	"os"
	"fmt"
)

var (
	exitCode int = 0
	c        chan os.Signal
	done     chan int
)

func main() {
	done = make(chan int)
	go ListenSignal()
	go Pull()
	go PrintLog()
	if _, ok := <-done; !ok {
		fmt.Println("Bye")
	}
	os.Exit(exitCode)
}
