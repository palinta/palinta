package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/matisszilard/palinta/palinta/store"
	"github.com/matisszilard/palinta/palinta/store/influxdb"
	"github.com/spf13/viper"
)

const (
	mqttBroker                     = "tcp://iot.eclipse.org:1883"
	influxDBHost                   = "localhost"
	influxDBPort                   = "8086"
	clientID                       = "palinta-go-client"
	mqttRootPath                   = "palinta"
	mqttRelayPath                  = "palinta/relay"
	mqttEnterLikeABossTopic        = "palinta/enterLikeABoss"
	mqttExitToken                  = "palinta/exit"
	mqttPrometheusTemperatureTopic = "palinta/prometheus/temperature"
)

var dbStore store.Store
var mqttChannel chan string

func main() {
	viper.SetConfigName("config.json")
	viper.AddConfigPath(".")

	viper.SetDefault("InfluxDBHost", "localhost")
	viper.SetDefault("InfluxDBPort", "8086")

	log.Info("Connect to Influx database...")
	dbStore = influxdb.New(influxDBHost, influxDBPort)
	log.Info("Connection created to Influx database!")

	mqttChannel = make(chan string, 2)

	log.Info("Connect to the MQTT broker")
	client, err := connectToMqttBroker(mqttBroker)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to connect to the MQTT broker")
		os.Exit(1)
	}
	err = subscribeForPalintaTopics(client)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to subscribe for tokens")
		os.Exit(1)
	}

	for {
		topic := <-mqttChannel
		payload := <-mqttChannel

		switch topic {
		case mqttEnterLikeABossTopic:
			{
				if payload == "start" {
					startImperialMarch()
				} else {
					stopImperialMarch()
				}
			}
		case mqttPrometheusTemperatureTopic:
			{
				saveTemperature(payload)
			}
		case mqttExitToken:
			{
				os.Exit(0)
			}
		}
	}
}

func subscribeForPalintaTopics(client mqtt.Client) error {
	err := subscribe(client, mqttEnterLikeABossTopic)
	if err != nil {
		return err
	}
	err = subscribe(client, mqttPrometheusTemperatureTopic)
	if err != nil {
		return err
	}
	err = subscribe(client, mqttExitToken)
	if err != nil {
		return err
	}
	return nil
}
