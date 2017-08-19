package g

import (
	"os"
	"sync"
	"github.com/lich2013/aria2cRPC"
)

var (
	ExitCode int             = 0
	C        chan os.Signal  = make(chan os.Signal)
	Done     chan int        = make(chan int)
	Conf     *Config         = new(Config)
	Wg       *sync.WaitGroup = new(sync.WaitGroup)
	TaskChan chan Task       = make(chan Task, 100)

	LogInfo      chan string    = make(chan string, 100)
	LogFatal     chan string    = make(chan string, 100)
	LogPath      string         = "/tmp/raspberry-client"
	LogInfoFile  string         = LogPath + "/raspberry-info.log"
	LogFatalFile string         = LogPath + "/raspberry-fatal.log"
	RPC          *aria2cRPC.RPC = aria2cRPC.RPC{}.Init("xiapian", "raspberry.lich.moe:6800")
)

type Config struct {
	Schema string `toml:"schema"`
	Host   string `toml:"host"`
}

type Task struct {
	TaskName   string `json:"task_name"`
	TaskType   string `json:"task_type"`
	TaskDetail string `json:"task_detail"`
	TaskId     string `json:"task_id"`
}
