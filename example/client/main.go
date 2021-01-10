package main

import (
	"context"
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/wilenceyao/humors"
	"github.com/wilenceyao/humors/example/api"
	"log"
	"time"
)

var client api.ExampleServiceClient

func main() {
	broker := flag.String("broker", "tcp://iot.eclipse.org:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	flag.Parse()
	clientid := "test2"
	mqttOpts := MQTT.NewClientOptions()
	mqttOpts.AddBroker(*broker)
	mqttOpts.SetClientID(clientid)
	mqttOpts.SetUsername(*user)
	mqttOpts.SetPassword(*password)
	rpcOpts := &humors.RPCOptions{
		Timeout: 2000,
	}
	opts := humors.Options{
		MQTTOpts: mqttOpts,
		RPCOpts:  rpcOpts,
	}
	h, err := humors.NewHumors(opts)
	if err != nil {
		log.Println("NewHumors err: ", err)
		return
	}
	// ms
	client = api.NewExampleServiceClient(h)
	// go executeLongTask()
	ticker := time.NewTicker(time.Millisecond * 500)
	var a int32 = 0
	for {
		<-ticker.C
		a = a + 1
		req := &api.SumRequest{
			A: a,
			B: 100,
		}
		res, err := client.Sum(context.Background(), "test1", req)
		if err != nil {
			log.Println("sum err: ", err)
		} else {
			log.Println("sum res: ", res)
		}

	}
}