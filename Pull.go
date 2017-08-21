package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"raspberry-client/g"
	"net/url"
	"strings"
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
	client := &http.Client{}
	req, err := http.NewRequest("GET", PullURL, nil)
	if err != nil {
		g.LogFatal <- err.Error()
	}
	req.Header.Add("Token", g.Conf.Token)
	for {
		resp, err := client.Do(req)
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
	client := &http.Client{}
	for x := range g.ReciveTaskChan {
		form := url.Values{"tasklist[]": []string{x.TaskId}}
		req, err := http.NewRequest("POST", ConfirmURL, strings.NewReader(form.Encode()))
		req.Header.Add("Token", g.Conf.Token)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
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
