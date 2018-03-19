package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/matisszilard/palinta/palinta/model"
	"github.com/matisszilard/palinta/palinta/store/influxdb"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	mqttBroker   = "tcp://iot.eclipse.org:1883"
	influxDBHost = "localhost"
	influxDBPort = "8086"
	//mqttBroker    = "tcp://192.168.99.247:1883"
	clientID                       = "palinta-go-client"
	mqttRootPath                   = "palinta"
	mqttRelayPath                  = mqttRootPath + "/relay"
	mqttEnterLikeABossTopic        = "palinta/enterLikeABoss"
	mqttExitToken                  = "palinta/exit"
	mqttPrometheusTemperatureTopic = "palinta/prometheus/temperature"
)

var exit = false

var mqttMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if msg.Topic() == mqttEnterLikeABossTopic {
		if string(msg.Payload()) == "start" {
			startImperialMarch()
		} else {
			stopImperialMarch()
		}
	}
	switch msg.Topic() {
	case mqttEnterLikeABossTopic:
		{
			if string(msg.Payload()) == "start" {
				startImperialMarch()
			} else {
				stopImperialMarch()
			}
		}
	case mqttPrometheusTemperatureTopic:
		{
			saveTemperature(string(msg.Payload()))
		}
	case mqttExitToken:
		{
			exit = true
		}

	}
}

func saveTemperature(temperature string) {
	var temp model.Temperature
	lt := time.Now()
	temp.Time = lt.String()

	t, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		t = 0.0
	}
	temp.Temperature = t

	store := influxdb.New(influxDBHost, influxDBPort)

	store.Temperatures().Save(temp)
}

func mqttSubscribe(topic string) {
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	opts.SetClientID(clientID)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(topic, 0, mqttMessageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		if !exit {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	c.Disconnect(250)
}

func sendMqttMessage(path string, message string) {
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	opts.SetClientID(clientID)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := c.Publish(path, 0, false, message)
	token.Wait()

	c.Disconnect(250)
}
