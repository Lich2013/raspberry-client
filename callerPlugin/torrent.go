package callerPlugin

import (
	"fmt"
	"raspberry-client/g"
)



func CallAria2c(x g.Task) {
	err := g.RPC.Ping()
	if err != nil {
		g.LogFatal <- fmt.Sprintf("task %+v faild, error is %s", x, err.Error())
		return
	}
	uri := []string{x.TaskDetail}
	ret, err := g.RPC.AddUri(uri)
	if err != nil {
		g.LogFatal <- fmt.Sprintf("task %+v faild, error is %s", x, err.Error())
		return
	}
	g.ReciveTaskChan <- x
	g.LogInfo <- fmt.Sprintf("task %+v success, task id is %s", x, ret)
}
