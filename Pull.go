package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"raspberry-client/g"
	"net/url"
)

var (
	host       string
	PullURL    string
	ConfirmURL string
)

type Result struct {
	Data   []g.Task        `json:"data"`
	Msg    string        `json:"msg"`
	Status int           `json:"status"`
}

func RegisterURL() {
	host = g.Conf.Host
	PullURL = host + "/tasklist"
	ConfirmURL = host + "/confirm"
}

func Pull() {
	RegisterURL()
	for {
		fmt.Println(time.Now())
		resp, err := http.Get(PullURL)
		if err != nil {
			g.LogFatal <- err.Error()
			time.Sleep(10 * time.Second)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			g.LogFatal <- err.Error()
			time.Sleep(10 * time.Second)
			continue
		}
		taskList := new(Result)
		json.Unmarshal(body, taskList)
		resp.Body.Close()
		for _, x := range taskList.Data {
			g.TaskChan <- x
		}
		time.Sleep(10 * time.Second)
	}
}

func Confirm() {
	for x := range g.ReciveTaskChan {
		resp, err := http.PostForm(ConfirmURL, url.Values{"tasklist[]": []string{x.TaskId}})
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			g.LogFatal <- err.Error()
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			g.LogFatal <- err.Error()
			continue
		}
		fmt.Println(string(body))
	}
	fmt.Println("confirm finish") // no possible to do
}
