package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var (
	host    string
	PullURL string
)

type Task struct {
	TaskName   string `json:"task_name"`
	TaskType   string `json:"task_type"`
	TaskDetail string `json:"task_detail"`
	TaskId     string `json:"task_id"`
}
type Result struct {
	Data   []Task        `json:"data"`
	Msg    string        `json:"msg"`
	Status int           `json:"status"`
}

func RegisterURL() {
	host = Conf.Schema + "://" + Conf.Host
	PullURL = host + "/tasklist"
}

func Pull() {
	RegisterURL()
	for {
		fmt.Println(time.Now())
		resp, err := http.Get(PullURL)
		if err != nil {
			LogFatal <- err.Error()
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			LogFatal <- err.Error()
		}
		taskList := new(Result)
		json.Unmarshal(body, taskList)
		resp.Body.Close()
		fmt.Println(taskList.Data)
		time.Sleep(10 * time.Second)
	}
}
