package queue

import (
	"defaultProject/components/config"

	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

type NsqQueue struct {
	MessageHandles nsq.Handler
}

func NewNsqQueue(handle nsq.Handler) *NsqQueue {
	return &NsqQueue{
		MessageHandles: handle,
	}
}

func (m *NsqQueue) Init() error {
	nsqHost := config.Cfg.GetString("Nsq.Host")
	channel := config.Cfg.GetString("Nsq.Channel")
	topic := config.Cfg.GetString("Nsq.Topic")

	poolSize := config.Cfg.GetInt("Nsq.HandleSize")
	cfg := nsq.NewConfig()
	for i := 0; i < poolSize; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
					panic(err)
				}
			}()

			c, err := nsq.NewConsumer(topic, channel, cfg)
			if err != nil {
				panic(err)
			}

			c.AddHandler(m.MessageHandles)

			if err := c.ConnectToNSQD(nsqHost); err != nil {
				panic(err)
			}
		}()
	}
	return nil
}
