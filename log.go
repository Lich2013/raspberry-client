package main

import (
	"log"
	"os"
	"syscall"
	"fmt"
)

var (
	LogInfo      chan string = make(chan string, 100)
	LogFatal     chan string = make(chan string, 100)
	LogPath      string      = "/tmp/raspberry-client"
	LogInfoFile  string      = LogPath + "/raspberry-info.log"
	LogFatalFile string      = LogPath + "/raspberry-fatal.log"
)

func PrintLog() {
	//CheckAndCreatFile()
	fd1, err := os.OpenFile(LogInfoFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	fd2, err := os.OpenFile(LogFatalFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	defer fd1.Close()
	defer fd2.Close()
	if err != nil {
		c <- syscall.SIGQUIT
	}
	info := log.New(fd1, "[INFO] ", log.LstdFlags)
	info.SetOutput(fd1)
	e := log.New(fd2, "[ERROR] ", log.LstdFlags)
	e.SetOutput(fd2)
	for {
		select {
		case loginfo := <-LogInfo:
			info.Println(loginfo)
		case logfatal := <-LogFatal:
			e.Println(logfatal)
		}
	}
}

func CheckAndCreatFile() {
	os.MkdirAll(LogPath, 0755)
	logList := []string{LogInfoFile, LogFatalFile}
	for _, x := range logList {
		_, err := os.Stat(x)
		if err != nil {
			f, err := os.Create(x)
			if err != nil {
				fmt.Println(err)
				c <- syscall.SIGQUIT
			}
			f.Close()
		}
	}

}
