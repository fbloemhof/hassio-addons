package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
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
	config := readConfig()

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
