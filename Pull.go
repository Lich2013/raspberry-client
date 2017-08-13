package main

import (
	"time"
	"fmt"
)

func Pull()  {
	for {
		fmt.Println(time.Now())
		time.Sleep(10 * time.Second)
	}
}
