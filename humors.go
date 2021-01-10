package humors

import (
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Humors struct {
	client   MQTT.Client
	Adaptors map[string]*HumorAdaptor

	Servants map[string]*HumorServant
	timeout  int32
}
type Options struct {
	MQTTOpts *MQTT.ClientOptions
	RPCOpts  *RPCOptions
}

type RPCOptions struct {
	// ms
	Timeout int32
}

var qos byte = 0

func formatAdaptorRecvResTopic(clientID, service string) string {
	return fmt.Sprintf("%s/%s/humorsres", clientID, service)
}

func formatServantRecvReqTopic(clientID, service string) string {
	return fmt.Sprintf("%s/%s/humorsreq", clientID, service)
}

func NewHumors(opts Options) (*Humors, error) {
	log.Println(DEBUG, "NewHumors")
	if opts.MQTTOpts == nil {
		return nil, errors.New("invalid opts")
	}
	c := MQTT.NewClient(opts.MQTTOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Println(ERROR, "connect mqtt err: %+v", token.Error())
		return nil, token.Error()
	}
	var rpcTimeout int32 = 2000
	if opts.RPCOpts != nil && opts.RPCOpts.Timeout != 0 {
		rpcTimeout = opts.RPCOpts.Timeout
	}
	h := &Humors{
		client:   c,
		Adaptors: make(map[string]*HumorAdaptor),
		Servants: make(map[string]*HumorServant),
		timeout:  rpcTimeout,
	}
	return h, nil
}

func (h *Humors) getClientOptionsReader() *MQTT.ClientOptionsReader {
	reader := h.client.OptionsReader()
	return &reader
}

func (h *Humors) NewAdaptor(service string) *HumorAdaptor {
	a := &HumorAdaptor{
		service: service,
		client:  h.client,
		timeout: h.timeout,
		topic:   formatAdaptorRecvResTopic(h.getClientOptionsReader().ClientID(), service),
	}
	h.Adaptors[service] = a
	a.init()
	return a
}

func (h *Humors) NewServant(service string) *HumorServant {
	topic := formatServantRecvReqTopic(h.getClientOptionsReader().ClientID(), service)
	s := &HumorServant{
		topic:   topic,
		client:  h.client,
		service: service,
	}
	h.Servants[service] = s
	s.init()
	return s
}

func publish(client MQTT.Client, topic string, btArr []byte) MQTT.Token {
	return client.Publish(topic, qos, false, btArr)
}
