package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"
	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
)

func init() {
	c = cron.New()

	c.AddFunc("@every 1m", detectExternelHearbeat)
	c.AddFunc("@every 1m", detectInternelHearbeat)
	c.AddFunc("@every 1m", fetchExternalNodes)
}

func getUrl(host string, port int, path string) string {
	return fmt.Sprintf("https://%s:%d%s", host, port, path)
}

func detectExternelHearbeat() {
	nodes := external.Record.GetNodes()
	for _, name := range nodes {
		resp, err := client.Get(getUrl(name, externalPort, "/healthz"))
		if err != nil {
			logger.Error(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			external.Record.RemoveByName(name)
		}
	}
}

func detectInternelHearbeat() {
	nodes := internal.Record.GetNodes()
	for _, name := range nodes {
		resp, err := client.Get(getUrl(name, internalPort, "/healthz"))
		if err != nil {
			logger.Error(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			internal.Record.RemoveByName(name)
		}
	}
}

func fetchExternalNodes() {
	nodes := external.Record.GetNodes()
	for _, name := range nodes {
		register(name, externalPort, externalHostName)
		externalMessage(name, externalPort, externalHostName)
	}
}

func externalMessage(host string, port int, local string) {
	logger.Infof("Post message from %s to: %s:%d", local, host, port)

	gwData, err := json.Marshal(deviceDataMap[externalHostName])
	if err != nil {
		logger.Error(err)
		return
	}

	request := entities.MessageRequest{
		Token: clientToken[host],
		Data:  string(gwData),
	}

	data, err := json.Marshal(request)
	if err != nil {
		logger.Error(err)
		return
	}

	resp, err := client.Post(getUrl(host, port, "/message"),
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		logger.Error(err)
		return
	}
	defer resp.Body.Close()

	response := &entities.MessageResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	json.Unmarshal(body, response)

	if resp.StatusCode != http.StatusOK {
		logger.Error(errors.New(response.Reason))
		return
	}

	if !response.Status {
		logger.Error(errors.New(response.Reason))
		return
	}
}

func register(host string, port int, local string) {
	request := &entities.RegisterRequest{
		Name: local,
	}

	data, err := json.Marshal(request)
	if err != nil {
		logger.Error(err)
		return
	}

	resp, err := client.Post(getUrl(host, port, "/register"),
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		logger.Error(err)
		return
	}
	defer resp.Body.Close()

	response := &entities.RegisterResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	json.Unmarshal(body, response)

	if resp.StatusCode != http.StatusOK {
		logger.Error(errors.New(response.Reason))
		return
	}

	if !response.Status {
		logger.Error(errors.New(response.Reason))
		return
	}

	clientToken[host] = response.Token
}
