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

func main() {
	broker := flag.String("broker", "tcp://iot.eclipse.org:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	flag.Parse()
	clientid := "test1"
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
	h.InitServant("")
	h.Servant.RegisterFun(1, api.SumRequest{}, api.SumResponse{}, sum)
	h.Servant.RegisterFun(2, api.SumRequest{}, api.SumResponse{}, longTask)
	ch := make(chan int, 1)
	<-ch
}

func sum(ctx context.Context, reqObj interface{}, resObj interface{}) {
	log.Println("sum rpc, req: ", reqObj)
	req := reqObj.(*api.SumRequest)
	res := resObj.(*api.SumResponse)
	res.Sum = req.A + req.B
	log.Println("sum, req: ", req, " res: ", res)
}

func longTask(ctx context.Context, reqObj interface{}, resObj interface{}) {
	log.Println("longTask rpc, req: ", reqObj)
	time.Sleep(time.Millisecond * 1500)
	req := reqObj.(*api.SumRequest)
	res := resObj.(*api.SumResponse)
	res.Sum = req.A + req.B
	log.Println("longTask, req: ", req, " res: ", res)
}
