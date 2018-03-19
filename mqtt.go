package main

import (
	log "github.com/Sirupsen/logrus"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.WithFields(log.Fields{
		"topic":   msg.Topic(),
		"payload": string(msg.Payload()),
	}).Info("Received an MQTT message")
	mqttChannel <- msg.Topic()
	mqttChannel <- string(msg.Payload())
}

func connectToMqttBroker(host string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	opts.SetClientID(clientID)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

func subscribe(client mqtt.Client, topic string) error {
	if token := client.Subscribe(topic, 0, mqttMessageHandler); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func publish(client mqtt.Client, topic string, message string) {
	token := client.Publish(topic, 0, false, message)
	token.Wait()
}
