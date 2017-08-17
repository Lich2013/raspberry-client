package main

import (
	"os"
	"fmt"
	"raspberry-client/g"
)

func InitAll() {
	go ListenSignal()
	go Pull()
	go PrintLog()
	go CallService()
}

func main() {
	g.Init()
	go InitAll()

	if _, ok := <-g.Done; !ok {
		fmt.Println("Bye")
	}
	os.Exit(g.ExitCode)
}
