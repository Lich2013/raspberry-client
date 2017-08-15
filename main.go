package main

import (
	"os"
	"fmt"
	"sync"
)

var (
	exitCode int             = 0
	c        chan os.Signal  = make(chan os.Signal)
	done     chan int        = make(chan int)
	Conf     *Config         = new(Config)
	wg       *sync.WaitGroup = new(sync.WaitGroup)
)

func InitAll() {
	go ListenSignal()
	go Pull()
	go PrintLog()
	go CallService()
}

func main() {
	wg.Add(1)
	InitConfig()
	wg.Wait()
	go InitAll()

	if _, ok := <-done; !ok {
		fmt.Println("Bye")
	}
	os.Exit(exitCode)
}
