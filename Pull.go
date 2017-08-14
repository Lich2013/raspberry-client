package main

import (
	"time"
	"fmt"
)

func Pull()  {
	for {
		fmt.Println(time.Now())
		LogInfo <- "6666"
		LogFatal <- "7777"
		time.Sleep(10 * time.Second)
	}
}
