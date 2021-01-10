package humors

import (
	"context"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type HumorAdaptor struct {
	// receive response topic
	topic  string
	client MQTT.Client
	reqID  int32
	resp   sync.Map
	// ms
	timeout int32
	service string
}

type ClientConnInterface interface {
	Invoke(ctx context.Context, topic string, method string, args interface{}, reply interface{}) error
}

func (a *HumorAdaptor) init() {
	a.resp = sync.Map{}
	a.client.Subscribe(a.topic, qos, a.MessageHandler)
}

func (a *HumorAdaptor) MessageHandler(_ MQTT.Client, message MQTT.Message) {
	if message.Topic() != a.topic {
		log.Println(ERROR, ADAPTOR, "unknown topic:", message.Topic())
	}
	resPkt := &ResponsePacket{}
	err := proto.Unmarshal(message.Payload(), resPkt)
	if err != nil {
		log.Println(ERROR, ADAPTOR, "unmarshal resPkt err:", err)
		return
	}
	ch, ok := a.resp.Load(resPkt.ReqID)
	if !ok {
		log.Println(WARN, ADAPTOR, "unknown resPkt reqID:", resPkt.ReqID)
		return
	}
	a.resp.Delete(resPkt.ReqID)
	ch.(chan *ResponsePacket) <- resPkt
}

func (a *HumorAdaptor) Call(ctx context.Context, clientID string, action string,
	req proto.Message, res proto.Message) error {
	topic := formatServantRecvReqTopic(clientID, a.service)
	begin := time.Now().UnixNano() / 1e6
	err := a.internalCallRpc(topic, action, req, res)
	end := time.Now().UnixNano() / 1e6
	rpc := fmt.Sprintf("%s/%s", a.service, action)
	log.Println(DEBUG, ADAPTOR, "rpc:", rpc, "reqTime:", begin,
		"resTime:", end, "cost:", end-begin)
	return err
}

func (a *HumorAdaptor) internalCallRpc(topic string, action string,
	req proto.Message, res proto.Message) error {
	reqID := atomic.AddInt32(&a.reqID, 1)
	reqBtArr, err := proto.Marshal(req)
	if err != nil {
		log.Println(ERROR, ADAPTOR, "proto marshal err:", err)
		return fmt.Errorf(HUMORS_ERR, ErrorCode_ENCODEERR)
	}
	reqPkt := &RequestPacket{
		ReqID:    reqID,
		Action:   action,
		Timeout:  a.timeout,
		Payload:  reqBtArr,
		ResTopic: a.topic,
	}
	readCh := make(chan *ResponsePacket)
	a.resp.Store(reqPkt.ReqID, readCh)
	reqPktBtArr, _ := proto.Marshal(reqPkt)
	token := publish(a.client, topic, reqPktBtArr)
	if !token.Wait() {
		log.Println(ERROR, ADAPTOR, "req push err:", token.Error())
		return fmt.Errorf(HUMORS_ERR, ErrorCode_CONNECTERR)
	}

	select {
	case <-time.After(time.Millisecond * time.Duration(a.timeout)):
		log.Println(ERROR, ADAPTOR, "request timeout")
		return fmt.Errorf(HUMORS_ERR, ErrorCode_TIMEOUT)
	case resPkt := <-readCh:
		if resPkt.Code != ErrorCode_SUCCESS {
			return fmt.Errorf(HUMORS_ERR, resPkt.Code)
		}
		return proto.Unmarshal(resPkt.Payload, res)
	}
}
