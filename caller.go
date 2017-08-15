package main

import (
	"fmt"
	"os/exec"
)

var (
	TaskChan chan Task = make(chan Task, 100)
)

func CallService() {
	for x := range TaskChan {
		switch (x.TaskType) {
		case "Torrent":
			go func() {
				cmd := exec.Command("ls")
				output, err := cmd.Output()
				fmt.Println(err, string(output))
				if err != nil {
					LogFatal <- fmt.Sprintf("task %+v faild, error is %s", x, err.Error())
					return
				}
				LogInfo <- fmt.Sprintf("task %+v success", x)
				fmt.Println(x)
			}()
		default:
			LogFatal <- fmt.Sprintf("unkonwn task %+v", x)
		}
	}

}
