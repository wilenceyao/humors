package humors

import (
	"context"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
	"sync"
	"time"
)

type MethodDesc struct {
	ReqType reflect.Type
	Fun     func(ctx context.Context, req proto.Message) (proto.Message, error)
}

type HumorServant struct {
	topic   string
	client  MQTT.Client
	funs    sync.Map
	service string
}

func (s *HumorServant) init() {
	s.funs = sync.Map{}
	s.client.Subscribe(s.topic, qos, s.MessageHandler)
}

func (s *HumorServant) RegisterMethod(action string, req interface{},
	fun func(ctx context.Context, req proto.Message) (proto.Message, error)) {
	log.Println(INFO, SERVANT, "register fun:", action)
	f := &MethodDesc{
		ReqType: reflect.TypeOf(req),
		Fun:     fun,
	}
	s.funs.Store(action, f)
}

func (s *HumorServant) MessageHandler(_ MQTT.Client, message MQTT.Message) {
	if message.Topic() != s.topic {
		log.Println(ERROR, SERVANT, "unknown topic:", message.Topic())
	}
	reqPkt := &RequestPacket{}
	err := proto.Unmarshal(message.Payload(), reqPkt)
	if err != nil {
		log.Println(ERROR, SERVANT, "unmarshal reqPkt err:", err)
		return
	}
	go s.executeRpc(reqPkt)
}

func (s *HumorServant) executeRpc(reqPkt *RequestPacket) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Millisecond*time.Duration(reqPkt.Timeout))
	defer cancel()
	resPkt := s.internalExecuteRpc(ctx, reqPkt)
	resPktBtArr, _ := proto.Marshal(resPkt)
	token := publish(s.client, reqPkt.ResTopic, resPktBtArr)
	if !token.Wait() {
		log.Println(ERROR, SERVANT, "res push err:", token.Error())
	}
}

func (s *HumorServant) internalExecuteRpc(ctx context.Context, reqPkt *RequestPacket) *ResponsePacket {
	resPkt := &ResponsePacket{
		ReqID: reqPkt.ReqID,
	}
	f, ok := s.funs.Load(reqPkt.Action)
	if !ok {
		resPkt.Code = ErrorCode_NOFUNCERR
		return resPkt
	}
	fun := f.(*MethodDesc)
	req := reflect.New(fun.ReqType).Interface().(proto.Message)
	err := proto.Unmarshal(reqPkt.Payload, req)
	if err != nil {
		log.Println(ERROR, SERVANT, "req payload unmarshal err:", err)
		resPkt.Code = ErrorCode_ENCODEERR
		return resPkt
	}
	begin := time.Now().UnixNano() / 1e6
	res, err := fun.Fun(ctx, req)
	end := time.Now().UnixNano() / 1e6
	rpc := fmt.Sprintf("%s/%s", s.service, reqPkt.Action)
	log.Println(DEBUG, SERVANT, "rpc:", rpc, "reqTime:", begin,
		"resTime:", end, "cost: ", end-begin)
	if err != nil {
		resPkt.Code = ErrorCode_SERVICEERR
		return resPkt
	}
	resBtArr, err := proto.Marshal(res.(proto.Message))
	if err != nil {
		resPkt.Code = ErrorCode_ENCODEERR
		return resPkt
	}
	resPkt.Payload = resBtArr
	return resPkt
}
