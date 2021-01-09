package humors

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Humors struct {
	client  MQTT.Client
	Adaptor *HumorAdaptor
	Servant *HumorServant
}

var qos byte = 0

func formatRecvResTopic(clientID string) string {
	return fmt.Sprintf("%s/humorsres", clientID)
}

func formatRecvReqTopic(clientID string) string {
	return fmt.Sprintf("%s/humorsreq", clientID)
}

func NewHumors(opts *MQTT.ClientOptions) (*Humors, error) {
	log.Println(DEBUG, "NewHumors")
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Println(ERROR, "connect mqtt err: %+v", token.Error())
		return nil, token.Error()
	}

	h := &Humors{
		client: c,
	}
	return h, nil
}

func (h *Humors) getClientOptionsReader() *MQTT.ClientOptionsReader {
	reader := h.client.OptionsReader()
	return &reader
}

func (h *Humors) InitAdaptor(timeout int32) {
	h.Adaptor = &HumorAdaptor{
		topic:   formatRecvResTopic(h.getClientOptionsReader().ClientID()),
		client:  h.client,
		timeout: timeout,
	}
	h.Adaptor.init()
}

func (h *Humors) InitServant(topic string) {
	if topic == "" {
		topic = formatRecvReqTopic(h.getClientOptionsReader().ClientID())
	}
	h.Servant = &HumorServant{
		topic:  topic,
		client: h.client,
	}
	h.Servant.init()
}

func publish(client MQTT.Client, topic string, btArr []byte) MQTT.Token {
	return client.Publish(topic, qos, false, btArr)
}
