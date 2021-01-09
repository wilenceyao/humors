package main

import (
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/wilenceyao/humors"
	"github.com/wilenceyao/humors/example/api"
	"log"
	"time"
)

func main() {
	broker := flag.String("broker", "tcp://iot.eclipse.org:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	flag.Parse()
	clientid := "test2"
	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(clientid)
	opts.SetUsername(*user)
	opts.SetPassword(*password)

	h, err := humors.NewHumors(opts)
	if err != nil {
		log.Println("NewHumors err: ", err)
		return
	}
	// ms
	h.InitAdaptor(2000)
	go executeLongTask(h)
	ticker := time.NewTicker(time.Millisecond * 100)
	var a int32 = 0
	for {
		<-ticker.C
		a = a + 1
		req := &api.SumRequest{
			A: a,
			B: 100,
		}
		res := &api.SumResponse{}
		err = h.Adaptor.Call("test1", 1, req, res)
		if err != nil {
			log.Println("sum err: ", err)
		} else {
			log.Println("sum res: ", res)
		}

	}
}

func executeLongTask(h *humors.Humors) {
	var err error
	ticker := time.NewTicker(time.Millisecond * 500)
	var a int32 = 0
	for {
		<-ticker.C
		a = a + 1
		req := &api.SumRequest{
			A: a,
			B: 100,
		}
		res := &api.SumResponse{}
		err = h.Adaptor.Call("test1", 2, req, res)
		if err != nil {
			log.Println("longtask err: ", err)
		} else {
			log.Println("longtask res: ", res)
		}

	}
}
