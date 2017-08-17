package main

import (
	"os/signal"
	"syscall"
	"fmt"
	"raspberry-client/g"
)

func ListenSignal()  {
	signal.Notify(g.C)
	for {
		sig := <-g.C
		fmt.Println(sig)
		g.ExitCode = 2
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			g.ExitCode = 0
		}
		if sig == syscall.SIGCHLD {
			continue
		}
		close(g.Done)
	}
}
