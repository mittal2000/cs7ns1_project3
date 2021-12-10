package main
// Swati Poojary
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"
	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
)

func init() {
	c = cron.New()

	c.AddFunc("@every 1m", fetchExternalNodes)
}

func getUrl(host string, port int, path string) string {
	return fmt.Sprintf("https://%s:%d%s", host, port, path)
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

	max := 100
	min := 10

	sensorData := &entities.SensorData{
		Name: externalHostName,
		Data: map[string]string{},
	}

	sensorData.Data["sensor1"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor2"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor3"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor4"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor5"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor6"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor7"] = strconv.Itoa(rand.Intn(max-min) + min)
	sensorData.Data["sensor8"] = strconv.Itoa(rand.Intn(max-min) + min)

	str, err := json.Marshal(sensorData)
	if err != nil {
		logger.Error(err)
	}

	ioutil.WriteFile(dir+"/data.json", str, 0644)

	gwData, err := json.Marshal(sensorData)
	if err != nil {
		logger.Error(err)
	}

	request := entities.MessageRequest{
		Token: clientToken[host],
		Data:  string(gwData),
	}

	data, err := json.Marshal(request)
	if err != nil {
		logger.Error(err)
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
