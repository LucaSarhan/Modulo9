package main

import (
	"testing"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)



func TestSubscriber(t *testing.T) {
	var messageArrived = make(chan struct{})

	var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		t.Logf("Received: %s\n", msg.Payload())
		close(messageArrived)
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	t.Log("Subscriber running.")

	select {
	case <-messageArrived:
		t.Log("Received message. Test passed.")
	}
}