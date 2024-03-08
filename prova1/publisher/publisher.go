package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fridgeorfreezer := [2]string{"fridge", "freezer"}
	index := 0

	for {
		location := fridgeorfreezer[index%2]
		text := Choice(location)
		token := client.Publish("Temperatura/topic", 0, false, text)
		token.Wait()
		fmt.Println(text)
		time.Sleep(2 * time.Second)
		index++
	}
}