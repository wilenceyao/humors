package main

import (
	"context"
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/wilenceyao/humors"
	"github.com/wilenceyao/humors/example/api"
	"log"
)

func main() {
	broker := flag.String("broker", "tcp://iot.eclipse.org:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	flag.Parse()
	clientid := "test1"
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
	impl := &ExampleImpl{}
	api.RegisterExampleServiceServer(h, impl)
	ch := make(chan int, 1)
	<-ch
}

type ExampleImpl struct {
}

func (i *ExampleImpl) Sum(ctx context.Context, req *api.SumRequest) (*api.SumResponse, error) {
	log.Println("sum rpc, req: ", req)
	res := &api.SumResponse{}
	res.Sum = req.A + req.B
	log.Println("sum, req: ", req, " res: ", res)
	return res, nil
}
