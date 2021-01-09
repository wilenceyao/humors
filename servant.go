package humors

import (
	"context"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
	"sync"
	"time"
)

type servantFun struct {
	ReqType reflect.Type
	ResType reflect.Type
	Fun     func(ctx context.Context, req interface{}, res interface{})
}

type HumorServant struct {
	topic  string
	client MQTT.Client
	funs   sync.Map
}

func (s *HumorServant) init() {
	s.funs = sync.Map{}
	// s.client.AddRoute(s.topic, s.MessageHandler)
	s.client.Subscribe(s.topic, qos, s.MessageHandler)
}

func (s *HumorServant) RegisterFun(action int32, req interface{}, res interface{},
	fun func(ctx context.Context, req interface{}, res interface{})) {
	log.Println(INFO, SERVANT, "register fun:", action)
	f := &servantFun{
		ReqType: reflect.TypeOf(req),
		ResType: reflect.TypeOf(res),
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
	log.Println(DEBUG, SERVANT, "recv req:", reqPkt.ReqID)
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
	fun := f.(*servantFun)
	req := reflect.New(fun.ReqType).Interface()
	res := reflect.New(fun.ResType).Interface()
	err := proto.Unmarshal(reqPkt.Payload, req.(proto.Message))
	if err != nil {
		log.Println(ERROR, SERVANT, "req payload unmarshal err:", err)
		resPkt.Code = ErrorCode_ENCODEERR
		return resPkt
	}
	begin := time.Now().UnixNano() / 1e6
	fun.Fun(ctx, req, res)
	end := time.Now().UnixNano() / 1e6
	log.Println(DEBUG, SERVANT, "rpc:", reqPkt.Action, "reqTime:", begin,
		"resTime:", end, "cost: ", end-begin)
	resBtArr, err := proto.Marshal(res.(proto.Message))
	if err != nil {
		resPkt.Code = ErrorCode_ENCODEERR
		return resPkt
	}
	resPkt.Payload = resBtArr
	return resPkt
}
