package main

import (
	"os/signal"
	"syscall"
	"fmt"
)

func ListenSignal()  {
	signal.Notify(c)
	for {
		sig := <-c
		fmt.Println(sig)
		exitCode = 2
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			exitCode = 0
		}
		close(done)
	}
}
