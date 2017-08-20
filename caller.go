package main

import (
	"fmt"
	"raspberry-client/callerPlugin"
	"raspberry-client/g"
)


func CallService() {
	for x := range g.TaskChan {
		switch (x.TaskType) {
		case "Torrent":
			go callerPlugin.CallAria2c(x)
		default:
			g.LogFatal <- fmt.Sprintf("unkonwn task %+v", x)
		}
	}

}


