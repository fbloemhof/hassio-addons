package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"net/http"
)

const configFileLocation = "/data/options.json"

type configType struct {
	AquareaServiceCloudURL      string
	AquareaServiceCloudLogin    string
	AquareaServiceCloudPassword string
	AquareaTimeout              string
	PoolInterval                string
	LogSecOffset                int64

	MqttServer    string
	MqttPort      int
	MqttLogin     string
	MqttPass      string
	MqttClientID  string
	MqttKeepalive string
}

type MQTTData struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func readConfig() configType {
	var config configType
	var configFile string
	configFile = configFileLocation

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting...")
	config := readConfig()

    // Set the API endpoint and the access token
    url := "http://supervisor/services/mqtt"
	req, err := http.NewRequest("GET", url, nil)
    supervisor_token, ok := os.LookupEnv("SUPERVISOR_TOKEN")
    if !ok {
        log.Fatalf("SUPERVISOR_TOKEN not set")
    }
    req.Header.Set("Authorization", "Bearer "+supervisor_token)
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error sending request: %s", err)
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        log.Fatalf("Error: %s", res.Status)
    }

    // Decode the response
    var data struct {
        Data  MQTTData `json:"data"`
    }
    if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
        log.Fatalf("Error decoding json: %s", err)
    }

    if config.MqttServer == "" {
        config.MqttServer = data.Data.Host
    }
    if config.MqttPort == 0 {
        config.MqttPort = data.Data.Port
    }
    if config.MqttLogin == "" {
        config.MqttLogin = data.Data.Username
    }
    if config.MqttPass == "" {
        config.MqttPass = data.Data.Password
    }

	dataChannel := make(chan map[string]string, 10)
	commandChannel := make(chan aquareaCommand, 10)
	statusChannel := make(chan bool) // offline-online

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)

	go mqttHandler(ctx, &wg, config, dataChannel, commandChannel, statusChannel)
	go aquareaHandler(ctx, &wg, config, dataChannel, commandChannel, statusChannel)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	log.Println("Shutting down")
	cancel()
	wg.Wait()
	log.Println("Shut down complete")
}
