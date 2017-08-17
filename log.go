package main

import (
	"log"
	"os"
	"syscall"
	"fmt"
	"raspberry-client/g"
)

func PrintLog() {
	//CheckAndCreatFile()
	fd1, err := os.OpenFile(g.LogInfoFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	fd2, err := os.OpenFile(g.LogFatalFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	defer fd1.Close()
	defer fd2.Close()
	if err != nil {
		g.C <- syscall.SIGQUIT
	}
	info := log.New(fd1, "[INFO] ", log.LstdFlags)
	info.SetOutput(fd1)
	e := log.New(fd2, "[ERROR] ", log.LstdFlags)
	e.SetOutput(fd2)
	for {
		select {
		case loginfo := <-g.LogInfo:
			info.Println(loginfo)
		case logfatal := <-g.LogFatal:
			e.Println(logfatal)
		}
	}
}

func CheckAndCreatFile() {
	os.MkdirAll(g.LogPath, 0755)
	logList := []string{g.LogInfoFile, g.LogFatalFile}
	for _, x := range logList {
		_, err := os.Stat(x)
		if err != nil {
			f, err := os.Create(x)
			if err != nil {
				fmt.Println(err)
				g.C <- syscall.SIGQUIT
			}
			f.Close()
		}
	}

}
