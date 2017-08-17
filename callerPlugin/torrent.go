package callerPlugin

import (
	"os/exec"
	"fmt"
	"raspberry-client/g"
)


func CallAria2c(x g.Task)  {
	cmd := exec.Command("ls")
	output, err := cmd.Output()
	fmt.Println(err, string(output))
	if err != nil {
		g.LogFatal <- fmt.Sprintf("task %+v faild, error is %s", x, err.Error())
		return
	}
	g.LogInfo <- fmt.Sprintf("task %+v success", x)
	fmt.Println(x)

}
