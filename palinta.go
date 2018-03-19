package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "relayOn",
			Usage: "Turns relay on",
			Action: func(c *cli.Context) error {
				sendMqttMessage(mqttRelayPath, "relayOn")
				return nil
			},
		},
		{
			Name:  "relayOff",
			Usage: "Turns relay off",
			Action: func(c *cli.Context) error {
				sendMqttMessage(mqttRelayPath, "relayOff")
				return nil
			},
		},
		{
			Name:  "enter-like-a-boss",
			Usage: "Play imperial march for entering like a boss",
			Action: func(c *cli.Context) error {
				mqttSubscribe(mqttEnterLikeABossTopic)
				return nil
			},
		},
		{
			Name:  "demeter",
			Usage: "Start demeter MQTT version",
			Action: func(c *cli.Context) error {
				mqttSubscribe(mqttPrometheusTemperatureTopic)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
